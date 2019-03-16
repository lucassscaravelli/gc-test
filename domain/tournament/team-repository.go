package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

// TeamRepository representa o repository de um torntimeeio
type TeamRepository struct {
	domaincore.IRepository
}

// NewRepository cria um novo repositorio
func NewTeamRepository() *TeamRepository {
	return &TeamRepository{gormrep.NewGormRepository()}
}

// FindAll retorna todos os time
func (s *TeamRepository) FindAll(where string, args ...interface{}) (m []*Team, err error) {
	m = []*Team{}
	err = s.IRepository.FindAll(&m, where, args...)
	return
}

// FindByID retorna um time buscando pelo ID
func (s *TeamRepository) FindByID(id uint) (t *Team, err error) {
	err = s.IRepository.FindByID(t, id)
	return
}

// Insert insere um time no banco de dados
func (s *TeamRepository) Insert(newTournament *Team) (t *Team, err error) {
	t = newTournament
	err = s.IRepository.Insert(t)
	return
}
