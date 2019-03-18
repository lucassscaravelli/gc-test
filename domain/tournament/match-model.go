package tournament

import (
	"gctest/errors"

	"github.com/jinzhu/gorm"
)

type Match struct {
	gorm.Model

	HostTeam      Team `gorm:"foreignkey:HostTeamRefer" json:"hostTeam"`
	HostTeamRefer uint

	VisitorTeam      Team `gorm:"foreignkey:VisitorTeamRefer" json:"visitorTeam"`
	VisitorTeamRefer uint

	HostScore    int `json:"hostScore"`
	VisitorScore int `json:"visitorScore"`

	GroupStage      GroupStage `gorm:"foreignkey:GroupStageRefer"`
	GroupStageRefer uint

	TeamBracket      TeamBracket `gorm:"foreignkey:TeamBracketRefer"`
	TeamBracketRefer uint
}

func (m *Match) GetID() uint {
	return m.ID
}

func (m *Match) Validate() error {

	if m.VisitorScore < 0 || m.HostScore < 0 {
		return errors.BadRequest
	}

	return nil
}

func (m *Match) GetWinner() (*Team, int) {
	if m.HostScore == 0 && m.VisitorScore == 0 {
		return nil, 0
	}

	if m.HostScore > m.VisitorScore {
		return &m.HostTeam, m.HostScore
	} else {
		return &m.VisitorTeam, m.VisitorScore
	}

}

func (m *Match) GetLoser() (*Team, int) {
	if m.HostScore == 0 && m.VisitorScore == 0 {
		return nil, 0
	}

	if m.HostScore < m.VisitorScore {
		return &m.HostTeam, m.HostScore
	} else {
		return &m.VisitorTeam, m.VisitorScore
	}

}
