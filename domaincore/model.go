package domaincore

// IModel é a estrutura generica de um modelo
type IModel interface {
	GetID() uint
	Validate() error
}
