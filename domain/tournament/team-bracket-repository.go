package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

type TeamBracketRepository struct {
	domaincore.IRepository
}

func NewTeamBracketRepository() *TeamBracketRepository {
	return &TeamBracketRepository{gormrep.NewGormRepository()}
}

func (s *TeamBracketRepository) Insert(newTeamBracket *TeamBracket) (m *TeamBracket, err error) {
	m = newTeamBracket
	err = s.IRepository.Insert(m)
	return
}

func (s *TeamBracketRepository) FindFirst(where string, args ...interface{}) (m *TeamBracket, err error) {
	m = &TeamBracket{}
	err = s.IRepository.FindFirst(m, where, args...)
	return
}
