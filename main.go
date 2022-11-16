package main

import (
	"github.com/ikiselewskii/avito-test-task/database"
	"github.com/ikiselewskii/avito-test-task/utils"
	"github.com/ikiselewskii/avito-test-task/webserver"
)

func main() {
	router := webserver.CreateRouterEngine()
	webserver.InitializeEndpoints(router)
	database.InitializeDBConnection(utils.SerializeDSN())
	database.CreateTables()
	router.Run()
}
