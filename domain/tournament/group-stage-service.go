package tournament

// GroupStageService representa o service de um torneio
type GroupStageService struct {
	rep *GroupStageRepository
}

// NewService cria um novo serviço de torneios
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
