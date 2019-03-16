package tournament

import (
	"gctest/errors"

	"github.com/jinzhu/gorm"
)

type Tournament struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t *Tournament) GetID() uint {
	return t.ID
}

func (t *Tournament) Validate() error {

	if t.Name == "" || t.Description == "" {
		return errors.BadRequest
	}

	return nil
}
