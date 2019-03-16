package database

import (
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"gctest/domain/tournament"
)

// Database - estrutura para gerenciar o banco de dados
type Database struct {
	Instance *gorm.DB
}

var models = []interface{}{(*tournament.Tournament)(nil), (*tournament.GroupStage)(nil), (*tournament.Team)(nil), (*tournament.TeamGroup)(nil), (*tournament.Match)(nil)}

// InitializePgDatabase inicializa o banco de dados, cria as tabelas
// e retorna um possivel erro
func InitializePgDatabase() (*Database, error) {
	var err error

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=testgc password=postgres")
	if err != nil {
		panic("failed to connect database")
	}

	dataStruct := Database{Instance: db}

	err = dataStruct.createTables()

	return &dataStruct, err
}

func (d *Database) createTables() error {

	for _, model := range models {
		err := d.Instance.CreateTable(model).Error
		if err != nil {

			if !strings.ContainsAny(err.Error(), "already exists") {
				return err
			}

		}
	}

	return nil
}
