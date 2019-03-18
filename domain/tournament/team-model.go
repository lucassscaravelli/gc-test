package tournament

import (
	"gctest/errors"

	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model

	Name  string `json:"name"`
	Tag   string `json:"tag"`
	Color string `json:"color"`

	TeamGroup []*TeamGroup `gorm:"many2many:group_teams"`
}

func (t *Team) GetID() uint {
	return t.ID
}

func (t *Team) Validate() error {

	if t.Name == "" || t.Tag == "" || t.Color == "" {
		return errors.BadRequest
	}

	return nil
}
