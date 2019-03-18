package gormrep

import (
	"gctest/domaincore"
	"gctest/errors"
	"gctest/logservice"
	"strings"

	"github.com/jinzhu/gorm"
)

type GormRepository struct {
	Db *gorm.DB
}

var dbInstance *gorm.DB
var logS = logservice.NewLogService("POSTGRESS - REPOSITORY")

func getInstance() *gorm.DB {
	return dbInstance
}

func InitializeRepository(Db *gorm.DB) {
	dbInstance = Db
}

func NewGormRepository() GormRepository {
	return GormRepository{Db: dbInstance}
}

func (gr GormRepository) GetDB() *gorm.DB {
	return gr.Db
}

func (gr GormRepository) Related(model interface{}, related interface{}, relatedTxt string) error {
	return gr.Db.Model(model).Related(related, relatedTxt).Error
}

func (gr GormRepository) Insert(model domaincore.IModel) error {
	if err := model.Validate(); err != nil {
		logS.Error(err.Error())
		return err
	}
	if err := gr.Db.Create(model).Error; err != nil {
		logS.Error(err.Error())
		return errors.ServerError
	}
	return nil
}

func (gr GormRepository) Update(model domaincore.IModel) error {
	if err := model.Validate(); err != nil {
		return err
	}
	return gr.Db.Save(model).Error
}

func (gr GormRepository) FindByID(receiver domaincore.IModel, id uint) (err error) {
	err = gr.Db.First(receiver, id).Error

	if strings.ContainsAny(err.Error(), "record not found") {
		return errors.NotFound
	}

	return
}

func (gr GormRepository) FindFirst(receiver domaincore.IModel, where string, args ...interface{}) error {
	err := gr.Db.Where(where, args...).Limit(1).Find(receiver).Error

	if err != nil {
		if strings.ContainsAny(err.Error(), "record not found") {
			receiver = nil
			return nil
		}
	}

	return err
}

func (gr GormRepository) FindAll(models interface{}, where string, args ...interface{}) (err error) {
	err = gr.Db.Where(where, args...).Find(models).Error
	return
}

func (gr GormRepository) Preload(model interface{}, column string, args ...interface{}) (err error) {
	err = gr.Db.Preload(column, args...).Find(model).Error
	return
}
