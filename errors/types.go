package errors

import (
	"errors"
)

var BadRequest = errors.New("Existem dados incorretos na requisição")
var ServerError = errors.New("Erro inesperado no servidor")
var NotFound = errors.New("Nenhum resultado encontrado")

var GroupStageAlreadyFinished = errors.New("A fase de grupos ja foi finalizada")
var PlayoffAlreadyFinished = errors.New("A fase de playoff ja foi finalizada")
var PreviousPhaseNotRun = errors.New("Para gerar essa fase é necessário gerar a fase anterior")

var InvalidPhase = errors.New("Fase não encontrada")
