package tournament

// MatchService representa o service de um torneio
type MatchService struct {
	rep *MatchRepository
}

// NewService cria um novo serviço de torneios
func NewMatchService() *MatchService {
	service := MatchService{NewMatchRepository()}

	return &service
}

func (s *MatchService) CreateMatch(newMatch *Match) (m *Match, err error) {
	m, err = s.rep.Insert(newMatch)
	return
}

func (s *MatchService) GetAllByGroupStageID(ID uint) (m []*Match, err error) {
	m, err = s.rep.FindAll("group_stage_refer = ?", ID)
	return
}

func (s *MatchService) UpdateMatch(match *Match) (m *Match, err error) {
	return s.rep.Update(match)
}
