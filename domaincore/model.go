package domaincore

type IModel interface {
	GetID() uint
	Validate() error
}
