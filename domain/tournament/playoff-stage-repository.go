package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

type PlayoffStageRepository struct {
	domaincore.IRepository
}

func NewPlayoffStageRepository() *PlayoffStageRepository {
	return &PlayoffStageRepository{gormrep.NewGormRepository()}
}

func (s *PlayoffStageRepository) Insert(newPlayoffStage *PlayoffStage) (m *PlayoffStage, err error) {
	m = newPlayoffStage
	err = s.IRepository.Insert(m)
	return
}
