package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

// TeamGroupRepository representa o repository de um torneio
type TeamGroupRepository struct {
	domaincore.IRepository
}

// NewRepository cria um novo repositorio
func NewTeamGroupRepository() *TeamGroupRepository {
	return &TeamGroupRepository{gormrep.NewGormRepository()}
}

// Insert insere um torneio no banco de dados
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
