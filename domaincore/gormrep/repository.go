package gormrep

import (
	"gctest/domaincore"
	"gctest/errors"
	"gctest/logservice"
	"strings"

	"github.com/jinzhu/gorm"
)

// GormRepository implementa a lib gorm
// para fazer operações no banco (ORM)
type GormRepository struct {
	Db *gorm.DB
}

var dbInstance *gorm.DB
var logS = logservice.NewLogService("POSTGRESS - REPOSITORY")

func getInstance() *gorm.DB {
	return dbInstance
}

// InitializeRepository inicializa os repositorios, ou seja,
// seta a instancia do banco de dados
func InitializeRepository(Db *gorm.DB) {
	dbInstance = Db
}

func NewGormRepository() GormRepository {
	return GormRepository{Db: dbInstance}
}

func (gr GormRepository) GetDB() *gorm.DB {
	return gr.Db
}

// func (gr *GormRepository) Initialize(args ...interface{}) {
// 	if len(args) == 0 {
// 		panic("*gorm.DB must be supplied for initialization")
// 	}
// 	if _, ok := args[0].(*gorm.DB); !ok {
// 		panic("The first arg must be *gorm.DB")
// 	}

// 	gr.InitDb(args[0].(*gorm.DB))
// }

// func (gr *GormRepository) InitDb(db *gorm.DB) {
// 	gr.Db = db
// }

func (gr GormRepository) Related(model interface{}, related interface{}, relatedTxt string) error {
	return gr.Db.Model(model).Related(related, relatedTxt).Error
}

// Insert insere um model genérico no banco
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

// func (gr GormRepository) Save(model IModel) (uint, error) {
// 	if err := model.Validate(); err != nil {
// 		return 0, err
// 	}
// 	if err := gr.Db.Save(model).Error; err != nil {
// 		return 0, err
// 	}
// 	return model.GetId(), nil
// }

func (gr GormRepository) FindByID(receiver domaincore.IModel, id uint) (err error) {
	err = gr.Db.First(receiver, id).Error

	if strings.ContainsAny(err.Error(), "record not found") {
		return errors.NotFound
	}

	return
}

func (gr GormRepository) FindFirst(receiver domaincore.IModel, where string, args ...interface{}) error {
	return gr.Db.Where(where, args...).Limit(1).Find(receiver).Error
}

func (gr GormRepository) FindAll(models interface{}, where string, args ...interface{}) (err error) {
	err = gr.Db.Where(where, args...).Find(models).Error
	return
}

func (gr GormRepository) Preload(model interface{}, column string, args ...interface{}) (err error) {
	err = gr.Db.Preload(column, args...).Find(model).Error
	return
}

// func (gr GormRepository) Delete(model IModel, where string, args ...interface{}) error {
// 	return gr.Db.Where(where, args...).Delete(&model).Error
// }

// func (gr GormRepository) NewRecord(model IModel) bool {
// 	return gr.Db.NewRecord(&model)
// }
