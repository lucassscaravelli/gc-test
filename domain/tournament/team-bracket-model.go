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

	// BracketNumber     int
	// PlayoffStageRefer uint

	// Match []*Match

	BracketType BracketType

	PlayoffStage      PlayoffStage `gorm:"foreignkey:PlayoffStageRefer"`
	PlayoffStageRefer uint
}

// GetID retorna o id do model
func (tg *TeamBracket) GetID() uint {
	return tg.ID
}

// Validate retorna um erro caso o schema
// não for válido
func (tg *TeamBracket) Validate() error {
	return nil
}
