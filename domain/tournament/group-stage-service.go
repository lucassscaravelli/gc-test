package tournament

import (
	"gctest/errors"
	"sort"
)

type GroupStageService struct {
	rep *GroupStageRepository
}

func NewGroupStageService() *GroupStageService {
	service := GroupStageService{NewGroupStageRepository()}

	return &service
}

func (s *GroupStageService) GetGroupStageByTournamentID(ID uint) (gs *GroupStage, err error) {
	gs, err = s.rep.FindFirst("tournament_refer = ?", ID)
	return
}

func (s *GroupStageService) CreateGroupStage(t *Tournament) (gs *GroupStage, err error) {

	newGroupStage := &GroupStage{Tournament: *t}
	gs, err = s.rep.Insert(newGroupStage)

	// obter os times
	teams, err := NewTeamService().GetTeamsForNewTournament(80)

	if err != nil {
		return nil, err
	}

	// criar grupos com os times obtidos
	groupStageGoups, err := NewTeamGroupService().CreateGroups(gs, teams)

	if err != nil {
		return nil, err
	}

	// criar partidas a partir dos grupos
	_, err = s.generateMatchesFromGroups(gs, groupStageGoups)

	return
}

func (s *GroupStageService) RunGroupStage(gs *GroupStage) (err error) {

	ms := NewMatchService()
	ps := NewPlayoffStageService()

	if playoffStage, err := ps.GetPlayoffStageByTournamentID(gs.TournamentRefer); err == nil && playoffStage.ID != 0 {
		return errors.GroupStageAlreadyFinished
	}

	matchesOfGs, err := ms.GetAllByGroupStageID(gs.ID)

	if err != nil {
		return err
	}

	for _, match := range matchesOfGs {
		_, err = ms.RunMatch(match)

		if err != nil {
			return err
		}
	}

	// a fase de grupo acabou, iniciar playoff
	_, err = ps.CreatePlayoffStage(&gs.Tournament)

	return
}

func (s *GroupStageService) GetTableResults(gs *GroupStage) (resultTable []*groupStageTable, err error) {

	// obter os grupos
	groups, err := NewTeamGroupService().GetTeamGroupsByGroupStageID(gs.GetID())

	if err != nil {
		return nil, err
	}

	// montar a estrutura de retornos
	for _, group := range groups {

		groupTeamList := []groupLine{}

		for _, team := range group.Teams {

			groupTeamList = append(groupTeamList, groupLine{
				TeamID:    team.ID,
				TeamColor: team.Color,
				TeamName:  team.Name,
				TeamTag:   team.Tag,
			})

		}

		resultTable = append(resultTable, &groupStageTable{GroupName: group.Group, Table: groupTeamList})

	}

	// obter as partidas
	matchesOfGs, err := NewMatchService().GetAllByGroupStageID(gs.ID)

	if err != nil {
		return nil, err
	}

	resultList := resultTable

	// popular a estrutura de retorno
	for _, match := range matchesOfGs {

		winnerTeam, winnerScore := match.GetWinner()
		loserTeam, loserScore := match.GetLoser()

		for i, groupTable := range resultTable {

			for j, line := range groupTable.Table {

				if winnerTeam != nil && line.TeamID == winnerTeam.ID {
					resultList[i].Table[j].TeamPoints++
					resultList[i].Table[j].TeamRoundBalance += (winnerScore - loserScore)

					resultList[i].Matches = append(resultList[i].Matches, matchInfo{
						HostID:    match.HostTeamRefer,
						HostScore: match.HostScore,
						HostColor: match.HostTeam.Color,
						HostTag:   match.HostTeam.Tag,

						VisitorID:    match.VisitorTeamRefer,
						VisitorScore: match.VisitorScore,
						VisitorColor: match.VisitorTeam.Color,
						VisitorTag:   match.VisitorTeam.Tag,
					})
				}

				if loserTeam != nil && line.TeamID == loserTeam.ID {
					resultList[i].Table[j].TeamRoundBalance += (loserScore - winnerScore)
				}

			}

		}
	}

	// ordenar
	for z := 0; z < len(resultList); z++ {

		sort.Slice(resultList[z].Table, func(i int, j int) bool {

			if resultList[z].Table[i].TeamPoints > resultList[z].Table[j].TeamPoints {
				return true
			} else if resultList[z].Table[i].TeamPoints < resultList[z].Table[j].TeamPoints {
				return false
			}

			if resultList[z].Table[i].TeamRoundBalance > resultList[z].Table[j].TeamRoundBalance {
				return true
			}

			return false

		})

	}

	return resultList, nil

}

func (s *GroupStageService) generateMatchesFromGroups(gs *GroupStage, teamGroups []*TeamGroup) (matches []*Match, err error) {

	for _, tg := range teamGroups {

		for i, teamH := range tg.Teams {

			for j := i + 1; j < len(tg.Teams); j++ {

				newMatch := Match{
					GroupStage:  *gs,
					HostTeam:    *teamH,
					VisitorTeam: *tg.Teams[j],
				}

				nM, err := NewMatchService().CreateMatch(&newMatch)

				if err != nil {
					return nil, err
				}

				matches = append(matches, nM)
			}

		}

	}

	return
}
