package tournament

import (
	"gctest/errors"
)

type PlayoffStageService struct {
	rep *PlayoffStageRepository
}

func NewPlayoffStageService() *PlayoffStageService {
	service := PlayoffStageService{NewPlayoffStageRepository()}

	return &service
}

func (s *PlayoffStageService) CreatePlayoffStage(t *Tournament) (m *PlayoffStage, err error) {

	teamService := NewTeamService()
	groupStageService := NewGroupStageService()
	teamBracketService := NewTeamBracketService()

	m = &PlayoffStage{Tournament: *t}
	m, err = s.rep.Insert(m)

	// Obter a fase de grupos do torneio
	groupStage, err := groupStageService.GetGroupStageByTournamentID(t.ID)

	if err != nil {

		return nil, err
	}

	// Obter a tabela com os dados formatados (grupos)
	tableResult, err := groupStageService.GetTableResults(groupStage)

	if err != nil {

		return nil, err
	}

	classifiedTeams := []*Team{}

	// Obter os dois primeiros colocados de cada grupo
	for _, group := range tableResult {

		firstTeam, err := teamService.GetTeamByID(group.Table[0].TeamID)

		if err != nil {
			return nil, err
		}

		secondTeam, err := teamService.GetTeamByID(group.Table[1].TeamID)

		if err != nil {

			return nil, err
		}

		classifiedTeams = append(classifiedTeams, firstTeam)
		classifiedTeams = append(classifiedTeams, secondTeam)

	}

	// gerar primeira fase
	_, err = teamBracketService.GenerateFirstPhase(m, classifiedTeams)

	return
}

func (s *PlayoffStageService) GetPlayoffStageByTournamentID(ID uint) (m *PlayoffStage, err error) {

	bracketService := NewTeamBracketService()

	m = &PlayoffStage{}
	m, err = s.rep.First("tournament_refer = ?", ID)

	if err != nil {
		return nil, err
	}

	m.FirstPhaseSeeds, err = bracketService.GetByPlayoffStageIDAndBracketType(m.ID, FirstPhase)

	if err != nil {
		return nil, err
	}

	m.OctavesSeeds, err = bracketService.GetByPlayoffStageIDAndBracketType(m.ID, OctavesPhase)

	if err != nil {
		return nil, err
	}

	m.Quartereeds, err = bracketService.GetByPlayoffStageIDAndBracketType(m.ID, QuarterPhase)

	if err != nil {
		return nil, err
	}

	m.SemiSeeds, err = bracketService.GetByPlayoffStageIDAndBracketType(m.ID, SemiPhase)

	if err != nil {
		return nil, err
	}

	m.FinalSeed, err = bracketService.GetByPlayoffStageIDAndBracketType(m.ID, Final)

	return
}

func (s *PlayoffStageService) GetTable(playoffStage *PlayoffStage) (table *playoffTable, err error) {

	table = &playoffTable{}

	if playoffStage.FirstPhaseSeeds.ID != 0 {
		table.FirstPhase, err = s.GetMatchStructFromPhase(playoffStage.FirstPhaseSeeds)
	}
	if playoffStage.OctavesSeeds.ID != 0 {
		table.Octaves, err = s.GetMatchStructFromPhase(playoffStage.OctavesSeeds)
	}
	if playoffStage.Quartereeds.ID != 0 {
		table.Quarter, err = s.GetMatchStructFromPhase(playoffStage.Quartereeds)
	}
	if playoffStage.SemiSeeds.ID != 0 {
		table.Semi, err = s.GetMatchStructFromPhase(playoffStage.SemiSeeds)
	}
	if playoffStage.FinalSeed.ID != 0 {
		table.Final, err = s.GetMatchStructFromPhase(playoffStage.FinalSeed)
	}

	return
}

func (s *PlayoffStageService) GetMatchStructFromPhase(bracket *TeamBracket) (matchStruct []matchInfo, err error) {

	matchService := NewMatchService()

	matches, err := matchService.GetAllByBracketID(bracket.ID)

	if err != nil {
		return nil, err
	}

	matchStruct = []matchInfo{}

	for _, match := range matches {

		matchStruct = append(matchStruct, matchInfo{
			HostColor: match.HostTeam.Color,
			HostID:    match.HostTeam.ID,
			HostScore: match.HostScore,
			HostTag:   match.HostTeam.Tag,

			VisitorColor: match.VisitorTeam.Color,
			VisitorID:    match.VisitorTeam.ID,
			VisitorScore: match.VisitorScore,
			VisitorTag:   match.VisitorTeam.Tag,
		})

	}

	if len(matchStruct) == 0 {
		return nil, nil
	}

	return

}

func (s *PlayoffStageService) RunNextPhase(playoffStage *PlayoffStage) (err error) {

	matchService := NewMatchService()

	// obter a tabela com dados formatados do playoff
	table, err := s.GetTable(playoffStage)

	// identificar a fase de grupo que esta
	if len(table.FirstPhase) == 0 {
		return errors.GroupStageHasNotInitiaized
	} else if len(table.Octaves) == 0 {
		err = s.runPhase(playoffStage, playoffStage.FirstPhaseSeeds, OctavesPhase)
	} else if len(table.Quarter) == 0 {
		err = s.runPhase(playoffStage, playoffStage.OctavesSeeds, QuarterPhase)
	} else if len(table.Semi) == 0 {
		err = s.runPhase(playoffStage, playoffStage.Quartereeds, SemiPhase)
	} else if len(table.Final) == 0 {
		err = s.runPhase(playoffStage, playoffStage.SemiSeeds, Final)
	} else {

		finalMatch, err := matchService.GetAllByBracketID(playoffStage.FinalSeed.ID)

		if err != nil {
			return err
		}

		// verificar se nao acabou
		if winner, _ := finalMatch[0].GetWinner(); winner == nil {
			return s.runPhase(playoffStage, playoffStage.FinalSeed, Nothing)
		}
		return errors.PlayoffAlreadyFinished

	}

	return
}

func (s *PlayoffStageService) runPhase(playoffStage *PlayoffStage, bracket *TeamBracket, nextBracketType BracketType) (err error) {

	matchService := NewMatchService()
	bracketService := NewTeamBracketService()

	// obter os jogos do bracket
	matches, err := matchService.GetAllByBracketID(bracket.ID)

	if err != nil {
		return err
	}

	// jogar todos os jogos
	for _, match := range matches {
		_, err = matchService.RunMatch(match)

		if err != nil {
			return err
		}

	}

	// gerar a proxima fase se nao
	// for a ultima fase
	if nextBracketType != Nothing {

		classifiedTeams, err := bracketService.getClassifiedTeams(bracket)

		if err != nil {
			return err
		}

		_, err = bracketService.getBracket(playoffStage, classifiedTeams, false, nextBracketType)

	}

	return
}
