package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
	"gctest/logservice"
)

var logS = logservice.NewLogService("DOMAIN.TOURNAMENT - REPOSITORY")

type TournamentRepository struct {
	domaincore.IRepository
}

func NewTournamentRepository() *TournamentRepository {
	return &TournamentRepository{gormrep.NewGormRepository()}
}

func (ts *TournamentRepository) FindAll() (t []*Tournament, err error) {
	err = ts.IRepository.FindAll(&t, "")
	return
}

// func (ts *TournamentRepository) FindByID(ID uint) (t *Tournament, err error) {
// 	err = ts.IRepository.FindByID(t, ID)
// 	return
// }

func (ts *TournamentRepository) FindFirst(where string, args ...interface{}) (t *Tournament, err error) {
	t = &Tournament{}
	err = ts.IRepository.FindFirst(t, where, args...)
	return
}

func (ts *TournamentRepository) Insert(newTournament *Tournament) (t *Tournament, err error) {
	t = newTournament
	err = ts.IRepository.Insert(t)
	return
}
