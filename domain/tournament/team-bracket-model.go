package tournament

import (
	"github.com/jinzhu/gorm"
)

type BracketType int

const (
	FirstPhase BracketType = iota
	OctavesPhase
	QuarterPhase
	SemiPhase
	Final
	Nothing
)

type TeamBracket struct {
	gorm.Model

	BracketType BracketType

	PlayoffStage      PlayoffStage `gorm:"foreignkey:PlayoffStageRefer"`
	PlayoffStageRefer uint
}

func (tg *TeamBracket) GetID() uint {
	return tg.ID
}
func (tg *TeamBracket) Validate() error {
	return nil
}
