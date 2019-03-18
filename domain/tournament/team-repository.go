package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

type TeamRepository struct {
	domaincore.IRepository
}

func NewTeamRepository() *TeamRepository {
	return &TeamRepository{gormrep.NewGormRepository()}
}

func (s *TeamRepository) FindAll(where string, args ...interface{}) (m []*Team, err error) {
	m = []*Team{}
	err = s.IRepository.FindAll(&m, where, args...)
	return
}

func (s *TeamRepository) FindByID(id uint) (t *Team, err error) {
	err = s.IRepository.FindByID(t, id)
	return
}

func (s *TeamRepository) Insert(newTournament *Team) (t *Team, err error) {
	t = newTournament
	err = s.IRepository.Insert(t)
	return
}

func (tr *TeamRepository) FindFirst(where string, args ...interface{}) (t *Team, err error) {
	t = &Team{}
	err = tr.IRepository.FindFirst(t, where, args...)
	return
}
