package tournament

// TeamGroupService representa o service de um torneio
type TeamGroupService struct {
	rep *TeamGroupRepository
}

// NewService cria um novo serviço de torneios
func NewTeamGroupService() *TeamGroupService {
	service := TeamGroupService{NewTeamGroupRepository()}

	return &service
}

func (s *TeamGroupService) GetTeamGroupsByGroupStageID(ID uint) (tg []*TeamGroup, err error) {
	tg, err = s.rep.FindAll("group_stage_refer = ?", ID)

	if err != nil {
		return nil, err
	}

	for _, teamGroup := range tg {
		err = s.rep.IRepository.Preload(teamGroup, "Teams")

		if err != nil {
			return nil, err
		}

	}

	return
}

func (s *TeamGroupService) CreateGroups(gs *GroupStage, teams []*Team) (tg []*TeamGroup, err error) {

	groupList := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}

	teamsInGroup := 0
	groupIndex := 0
	groupTeams := []*Team{}

	for _, teamItem := range teams {

		if teamsInGroup == 5 {

			newTeamGroup := &TeamGroup{GroupStage: *gs, Group: groupList[groupIndex], Teams: groupTeams}
			newTeamGroup, err = s.rep.Insert(newTeamGroup)

			if err != nil {
				return nil, err
			}

			tg = append(tg, newTeamGroup)

			teamsInGroup = 0
			groupIndex++

			groupTeams = []*Team{}
		}

		groupTeams = append(groupTeams, teamItem)
		teamsInGroup++
	}

	// inserir o ultimo grupo
	newTeamGroup := &TeamGroup{GroupStage: *gs, Group: groupList[groupIndex], Teams: groupTeams}
	newTeamGroup, err = s.rep.Insert(newTeamGroup)

	if err != nil {
		return nil, err
	}

	tg = append(tg, newTeamGroup)

	return
}
