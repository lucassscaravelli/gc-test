package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

// GroupStageRepository representa o repository de um group-stage
type GroupStageRepository struct {
	domaincore.IRepository
}

// NewRepository cria um novo repositorio
func NewGroupStageRepository() *GroupStageRepository {
	return &GroupStageRepository{gormrep.NewGormRepository()}
}

// FindByID retorna um group-stage buscando pelo ID
func (gtr *GroupStageRepository) FindByID(id uint) (gs *GroupStage, err error) {
	err = gtr.IRepository.FindByID(gs, id)
	return
}

// FindByID retorna um group-stage buscando pelo ID
func (gtr *GroupStageRepository) FindFirst(where string, args ...interface{}) (gs *GroupStage, err error) {
	gs = &GroupStage{}
	err = gtr.IRepository.FindFirst(gs, where, args...)
	return
}

// Insert insere um group-stage no banco de dados
func (gtr *GroupStageRepository) Insert(newGs *GroupStage) (gs *GroupStage, err error) {
	gs = newGs
	err = gtr.IRepository.Insert(gs)
	return
}
