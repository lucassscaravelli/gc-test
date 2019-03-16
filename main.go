package main

import (
	"gctest/api"
	"gctest/database"
	"gctest/domaincore/gormrep"
	"gctest/logservice"
)

var logS = logservice.NewLogService("main")

func main() {

	logS.Info("Criando conexão com banco de dados")
	db, err := database.InitializePgDatabase()
	defer db.Instance.Close()

	if err != nil {
		panic(err)
	}

	logS.Info("Inicializando repositorios")
	gormrep.InitializeRepository(db.Instance)

	logS.Info("Inicializando server")
	api.RunServer()
}
