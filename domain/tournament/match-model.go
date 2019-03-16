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
}

// GetID retorna o id do model
func (m *Match) GetID() uint {
	return m.ID
}

// Validate retorna um erro caso o schema
// não for válido
func (m *Match) Validate() error {

	if m.VisitorScore < 0 || m.HostScore < 0 {
		return errors.BadRequest
	}

	return nil
}

func (m *Match) GetWinner() (uint, int) {
	if m.HostScore == 0 && m.VisitorScore == 0 {
		return 0, 0
	}

	if m.HostScore > m.VisitorScore {
		return m.HostTeamRefer, m.HostScore
	} else {
		return m.VisitorTeamRefer, m.VisitorScore
	}

}

func (m *Match) GetLoser() (uint, int) {
	if m.HostScore == 0 && m.VisitorScore == 0 {
		return 0, 0
	}

	if m.HostScore < m.VisitorScore {
		return m.HostTeamRefer, m.HostScore
	} else {
		return m.VisitorTeamRefer, m.VisitorScore
	}

}
