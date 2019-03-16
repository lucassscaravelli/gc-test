package tournament

type TournamentService struct {
	rep *TournamentRepository
}

func NewTournamentService() *TournamentService {
	service := TournamentService{NewTournamentRepository()}

	return &service
}

func (s *TournamentService) GetAllTournaments() (t []*Tournament, err error) {
	t, err = s.rep.FindAll()
	return
}

func (s *TournamentService) GetTournamentByID(ID uint) (t *Tournament, err error) {
	t, err = s.rep.FindFirst("id = ?", ID)
	return
}

func (s *TournamentService) CreateTournament(newTournament *Tournament) (t *Tournament, err error) {

	t, err = s.rep.Insert(newTournament)

	if err != nil {
		return nil, err
	}

	_, err = NewGroupStageService().CreateGroupStage(t)

	if err != nil {
		return nil, err
	}

	return
}
