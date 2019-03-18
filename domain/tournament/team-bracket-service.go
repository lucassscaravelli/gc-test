package tournament

type TeamBracketService struct {
	rep *TeamBracketRepository
}

func NewTeamBracketService() *TeamBracketService {
	service := TeamBracketService{NewTeamBracketRepository()}

	return &service
}

func (tbs *TeamBracketService) GenerateFirstPhase(playoffStage *PlayoffStage, teams []*Team) (bracket *TeamBracket, err error) {
	bracket, err = tbs.getBracket(playoffStage, teams, true, FirstPhase)
	return
}

func (tbs *TeamBracketService) GetByPlayoffStageIDAndBracketType(ID uint, bracketType BracketType) (bracket *TeamBracket, err error) {
	bracket, err = tbs.rep.FindFirst("playoff_stage_refer = ? AND bracket_type = ?", ID, bracketType)
	return
}

func (tbs *TeamBracketService) getClassifiedTeams(bracket *TeamBracket) (teams []*Team, err error) {
	ms := NewMatchService()
	teams = []*Team{}

	matches, err := ms.GetAllByBracketID(bracket.ID)

	if err != nil {
		return nil, err
	}

	teams = []*Team{}

	for _, match := range matches {

		winner, _ := match.GetWinner()

		teams = append(teams, winner)
	}

	return
}

func (tbs *TeamBracketService) getBracket(playOffStage *PlayoffStage, teams []*Team, alterSeed bool, bracketType BracketType) (bracket *TeamBracket, err error) {

	matchService := NewMatchService()

	teamsLen := len(teams)
	matchersOfBracketLen := teamsLen / 2
	hostIndex := 0

	var visitorIndex int

	if alterSeed == true {
		visitorIndex = teamsLen - 1
	} else {
		visitorIndex = 1
	}

	bracket = &TeamBracket{
		BracketType:  bracketType,
		PlayoffStage: *playOffStage,
	}

	bracket, err = tbs.rep.Insert(bracket)

	if err != nil {
		return nil, err
	}

	for i := 1; i <= matchersOfBracketLen; i++ {

		matchOfBucket := &Match{
			TeamBracket: *bracket,
			HostTeam:    *teams[hostIndex],
			VisitorTeam: *teams[visitorIndex],
		}

		if _, err := matchService.CreateMatch(matchOfBucket); err != nil {
			return nil, err
		}

		if alterSeed == true {
			hostIndex++
			visitorIndex--
		} else {
			hostIndex += 2
			visitorIndex = hostIndex + 1
		}

	}

	return
}

func (tbs *TeamBracketService) GetMatchFromBracket(bracket *TeamBracket) (match *Match, err error) {

	if bracket == nil {
		return nil, nil
	}

	ms := NewMatchService()
	match, err = ms.GetMatchByBracketID(bracket.ID)
	return
}
