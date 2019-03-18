package tournament

import (
	"gctest/domaincore"
	"gctest/domaincore/gormrep"
)

type GroupStageRepository struct {
	domaincore.IRepository
}

func NewGroupStageRepository() *GroupStageRepository {
	return &GroupStageRepository{gormrep.NewGormRepository()}
}

func (gtr *GroupStageRepository) FindByID(id uint) (gs *GroupStage, err error) {
	err = gtr.IRepository.FindByID(gs, id)
	return
}

func (gtr *GroupStageRepository) FindFirst(where string, args ...interface{}) (gs *GroupStage, err error) {
	gs = &GroupStage{}
	err = gtr.IRepository.FindFirst(gs, where, args...)

	if err != nil {
		return nil, err
	}

	err = gtr.IRepository.Preload(gs, "Tournament")

	return
}

func (gtr *GroupStageRepository) Insert(newGs *GroupStage) (gs *GroupStage, err error) {
	gs = newGs
	err = gtr.IRepository.Insert(gs)
	return
}
