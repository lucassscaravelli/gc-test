package tournament

import (
	"github.com/jinzhu/gorm"
)

type PlayoffStage struct {
	gorm.Model

	Tournament      Tournament `gorm:"foreignkey:TournamentRefer"`
	TournamentRefer uint
}

func (gs *PlayoffStage) GetID() uint {
	return gs.ID
}

func (gs *PlayoffStage) Validate() error {
	return nil
}
