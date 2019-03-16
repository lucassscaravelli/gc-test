package errors

import (
	"errors"
)

var BadRequest = errors.New("Existem dados incorretos na requisição")
var ServerError = errors.New("Erro inesperado no servidor")
var NotFound = errors.New("Nenhum resultado encontrado")