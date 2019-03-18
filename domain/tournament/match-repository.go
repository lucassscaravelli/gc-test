package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

type MatchRepository struct {
	domaincore.IRepository
}

func NewMatchRepository() *MatchRepository {
	return &MatchRepository{gormrep.NewGormRepository()}
}

func (s *MatchRepository) Insert(newMatch *Match) (m *Match, err error) {
	m = newMatch
	err = s.IRepository.Insert(m)
	return
}

func (ts *MatchRepository) FindAll(where string, args ...interface{}) (m []*Match, err error) {
	m = []*Match{}
	err = ts.IRepository.FindAll(&m, where, args...)
	return
}

func (ts *MatchRepository) Update(match *Match) (m *Match, err error) {
	m = match
	err = ts.IRepository.Update(m)
	return
}

func (ts *MatchRepository) FindFirst(where string, args ...interface{}) (m *Match, err error) {
	m = &Match{}
	err = ts.IRepository.FindFirst(m, where, args...)
	return
}
