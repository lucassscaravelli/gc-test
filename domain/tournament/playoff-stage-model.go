package tournament

import (
	"github.com/jinzhu/gorm"
)

type PlayoffStage struct {
	gorm.Model

	Tournament      Tournament `gorm:"foreignkey:TournamentRefer"`
	TournamentRefer uint

	FirstPhaseSeeds *TeamBracket
	OctavesSeeds    *TeamBracket
	Quartereeds     *TeamBracket
	SemiSeeds       *TeamBracket
	FinalSeed       *TeamBracket
}

type playoffTable struct {
	FirstPhase []matchInfo
	Octaves    []matchInfo
	Quarter    []matchInfo
	Semi       []matchInfo
	Final      []matchInfo
}

func (gs *PlayoffStage) GetID() uint {
	return gs.ID
}

func (gs *PlayoffStage) Validate() error {
	return nil
}
