package tournament

import (
	"fmt"
	"math/rand"
	"strings"
)

// TeamService representa o service de um time
type TeamService struct {
	rep *TeamRepository
}

var prefixNameList = []string{
	"The",
	"Angry",
	"Kings Of",
	"Sons Of",
	"Lords Of",
}

var nameList = []string{
	"Kings",
	"Lords",
	"Nobles",
	"Gods",
	"Odin",
	"Marvel",
	"Dust_2",
}

var sufixNameLIst = []string{
	"Squad",
	"Team",
	"Org",
}

// NewService cria um novo serviço de times
func NewTeamService() *TeamService {
	service := TeamService{NewTeamRepository()}

	return &service
}

// GetAllTeams retorna todos os times
func (s *TeamService) GetAllTeams() (t []*Team, err error) {
	t, err = s.rep.FindAll("")
	return
}

func (s *TeamService) GetTeamsByGroupID(ID uint) (t []*Team, err error) {
	t, err = s.rep.FindAll("team_group_id = ?", ID)
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
		tName := getRandomName()
		tTag := strings.ToUpper(tName[0:3])
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

// CreateTeam cria um time
func (s *TeamService) CreateTeam(newTeam *Team) (t *Team, err error) {
	t, err = s.rep.Insert(newTeam)
	return
}

func getRandomName() string {
	iPrefix := rand.Intn(len(prefixNameList))
	iName := rand.Intn(len(nameList))
	iSufix := rand.Intn(len(sufixNameLIst))

	return fmt.Sprintf("%s %s %s", prefixNameList[iPrefix], nameList[iName], sufixNameLIst[iSufix])
}

func getRandomColor() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}
