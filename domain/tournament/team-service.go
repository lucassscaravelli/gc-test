package tournament

import (
	"fmt"
	"math/rand"
)

type TeamService struct {
	rep *TeamRepository
}

var prefixNameList = []string{
	"The",
	"Angry",
	"Kings Of",
	"Sons Of",
	"Lords Of",
	"Masters of",
	"Brothers of",
	"Happy",
}

var nameList = []string{
	"Kings",
	"Lords",
	"Nobles",
	"Gods",
	"Odin",
	"Marvel",
	"Dust_2",
	"Brazil",
	"America",
	"Aliens",
}

var sufixNameLIst = []string{
	"Kingdom",
	"Alliance",
	"Initiative",
	"Squad",
	"Team",
	"Org",
}

func NewTeamService() *TeamService {
	service := TeamService{NewTeamRepository()}

	return &service
}

func (s *TeamService) GetAllTeams() (t []*Team, err error) {
	t, err = s.rep.FindAll("")
	return
}

func (s *TeamService) GetTeamsByGroupID(ID uint) (t []*Team, err error) {
	t, err = s.rep.FindAll("team_group_id = ?", ID)
	return
}

func (s *TeamService) GetTeamByID(ID uint) (t *Team, err error) {
	t = &Team{}
	t, err = s.rep.FindFirst("id = ?", ID)
	return
}

func (s *TeamService) GetTeamsForNewTournament(count int) (t []*Team, err error) {
	t, err = s.GetAllTeams()

	if len(t) < count {
		if err := s.generateRandomTeams(count); err != nil {
			return nil, err
		}

		t, err = s.GetAllTeams()
	}

	return
}

func (s *TeamService) generateRandomTeams(count int) error {
	for i := 0; i < count; i++ {
		tName, tTag := getRandomNameAndTag()
		tColor := getRandomColor()

		newTeam := &Team{
			Name:  tName,
			Tag:   tTag,
			Color: tColor,
		}

		if _, err := s.CreateTeam(newTeam); err != nil {
			return err
		}
	}

	return nil
}

func (s *TeamService) CreateTeam(newTeam *Team) (t *Team, err error) {
	t, err = s.rep.Insert(newTeam)
	return
}

func getRandomNameAndTag() (string, string) {
	iPrefix := rand.Intn(len(prefixNameList))
	iName := rand.Intn(len(nameList))
	iSufix := rand.Intn(len(sufixNameLIst))

	return fmt.Sprintf("%s %s %s", prefixNameList[iPrefix], nameList[iName], sufixNameLIst[iSufix]),
		fmt.Sprintf("%c%c%c", prefixNameList[iPrefix][0], nameList[iName][0], sufixNameLIst[iSufix][0])
}

func getRandomColor() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}
