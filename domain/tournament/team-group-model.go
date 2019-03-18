package tournament

import (
	"gctest/errors"

	"github.com/jinzhu/gorm"
)

type TeamGroup struct {
	gorm.Model

	Group string `json:"group"`

	GroupStage      GroupStage `gorm:"foreignkey:GroupStageRefer"`
	GroupStageRefer uint

	Teams []*Team `gorm:"many2many:group_teams" json:"teams"`
}

func (tg *TeamGroup) GetID() uint {
	return tg.ID
}

func (tg *TeamGroup) Validate() error {

	if tg.Group == "" || tg.Teams == nil {
		return errors.BadRequest
	}

	return nil
}
