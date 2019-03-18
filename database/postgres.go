package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"gctest/domain/tournament"
)

type Database struct {
	Instance *gorm.DB
}

var models = []interface{}{(*tournament.TeamBracket)(nil),
	(*tournament.PlayoffStage)(nil), (*tournament.Tournament)(nil),
	(*tournament.GroupStage)(nil), (*tournament.Team)(nil),
	(*tournament.TeamGroup)(nil), (*tournament.Match)(nil)}

func InitializePgDatabase() (*Database, error) {
	var err error
	var db *gorm.DB

	first := true

	for first || err != nil {

		db, err = gorm.Open("postgres", "host=postgres_db port=5432 user=postgres dbname=testgc password=postgres sslmode=disable")

		if err != nil {
			fmt.Println(err.Error())
		}

		time.Sleep(5 * time.Second)

		first = false
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
