package tournament

// TournamentService representa o service de um torneio
type TournamentService struct {
	rep *TournamentRepository
}

// NewService cria um novo serviço de torneios
func NewTournamentService() *TournamentService {
	service := TournamentService{NewTournamentRepository()}

	return &service
}

// GetAllTournaments retorna todos os torneios
func (s *TournamentService) GetAllTournaments() (t []*Tournament, err error) {
	t, err = s.rep.FindAll()
	return
}

// GetTournamentByID retorna um torneio a partir de um id
func (s *TournamentService) GetTournamentByID(ID uint) (t *Tournament, err error) {
	t, err = s.rep.FindFirst("id = ?", ID)
	return
}

// CreateTournament cria um torneio
func (s *TournamentService) CreateTournament(newTournament *Tournament) (t *Tournament, err error) {

	t, err = s.rep.Insert(newTournament)

	if err != nil {
		return nil, err
	}

	_, err = NewGroupStageService().CreateGroupStage(t)

	return
}
