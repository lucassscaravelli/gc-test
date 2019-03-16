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

// GetID retorna o id do model
func (tg *TeamGroup) GetID() uint {
	return tg.ID
}

// Validate retorna um erro caso o schema
// não for válido
func (tg *TeamGroup) Validate() error {

	if tg.Group == "" || tg.Teams == nil {
		return errors.BadRequest
	}

	return nil
}
