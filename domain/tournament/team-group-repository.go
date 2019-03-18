package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

type TeamGroupRepository struct {
	domaincore.IRepository
}

func NewTeamGroupRepository() *TeamGroupRepository {
	return &TeamGroupRepository{gormrep.NewGormRepository()}
}

func (s *TeamGroupRepository) Insert(newTg *TeamGroup) (tg *TeamGroup, err error) {
	tg = newTg
	err = s.IRepository.Insert(tg)
	return
}

func (s *TeamGroupRepository) FindAll(where string, args ...interface{}) (m []*TeamGroup, err error) {
	m = []*TeamGroup{}
	err = s.IRepository.FindAll(&m, where, args...)
	return
}
