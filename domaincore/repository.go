package domaincore

// IRepository é a estrutura de um rep. generico
type IRepository interface {
	Insert(model IModel) error

	Update(model IModel) error

	Related(model interface{}, related interface{}, relatedTxt string) error

	Preload(model interface{}, column string, args ...interface{}) error

	FindByID(receiver IModel, ID uint) error

	FindFirst(receiver IModel, where string, args ...interface{}) error

	FindAll(models interface{}, where string, args ...interface{}) (err error)
}
