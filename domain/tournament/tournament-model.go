package tournament

import (
	"gctest/errors"

	"github.com/jinzhu/gorm"
)

// Tournament é o modelo que representa um torneio
type Tournament struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetID retorna o id do model
func (t *Tournament) GetID() uint {
	return t.ID
}

// Validate retorna um erro caso o schema
// não for válido
func (t *Tournament) Validate() error {

	if t.Name == "" || t.Description == "" {
		return errors.BadRequest
	}

	return nil
}

func (t *Tournament) GetGroupStage() (gs *GroupStage, err error) {
	return
}
