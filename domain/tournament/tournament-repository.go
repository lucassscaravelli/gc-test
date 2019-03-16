package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

type TournamentRepository struct {
	domaincore.IRepository
}

func NewTournamentRepository() *TournamentRepository {
	return &TournamentRepository{gormrep.NewGormRepository()}
}

func (tr *TournamentRepository) FindAll() (t []*Tournament, err error) {
	err = tr.IRepository.FindAll(&t, "")
	return
}

func (tr *TournamentRepository) FindFirst(where string, args ...interface{}) (t *Tournament, err error) {
	t = &Tournament{}
	err = tr.IRepository.FindFirst(t, where, args...)
	return
}

func (tr *TournamentRepository) Insert(newTournament *Tournament) (t *Tournament, err error) {
	t = newTournament
	err = tr.IRepository.Insert(t)
	return
}
