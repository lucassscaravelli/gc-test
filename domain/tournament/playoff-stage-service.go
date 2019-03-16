package tournament

type PlayoffStageService struct {
	rep *PlayoffStageRepository
}

func NewPlayoffStageService() *PlayoffStageService {
	service := PlayoffStageService{NewPlayoffStageRepository()}

	return &service
}

func (s *PlayoffStageService) CreatePlayoffStage(t *Tournament) (m *PlayoffStage, err error) {
	m = &PlayoffStage{Tournament: *t}
	m, err = s.rep.Insert(m)
	return
}
