package tournament

import (
	"github.com/jinzhu/gorm"
)

type GroupStage struct {
	gorm.Model

	Tournament      Tournament `gorm:"foreignkey:TournamentRefer"`
	TournamentRefer uint
}

type groupStageTable struct {
	GroupName string
	Table     []groupLine
	Matches   []matchInfo
}

type groupLine struct {
	TeamID           uint
	TeamColor        string
	TeamName         string
	TeamTag          string
	TeamRoundBalance int
	TeamPoints       int
}

type matchInfo struct {
	HostID    uint
	HostTag   string
	HostColor string
	HostScore int

	VisitorID    uint
	VisitorTag   string
	VisitorColor string
	VisitorScore int
}

func (gs *GroupStage) GetID() uint {
	return gs.ID
}

func (gs *GroupStage) Validate() error {
	return nil
}
