package main

import (
	"github.com/ikiselewskii/avito-test-task/database"
	"github.com/ikiselewskii/avito-test-task/utils"
	"github.com/ikiselewskii/avito-test-task/webserver"
)

func main() {
	router := webserver.CreateRouterEngine()
	webserver.InitializeEndpoints(router)
	db := database.InitializeDBConnection(utils.ParseEnvToDSN())
	database.CreateTables(db)
	router.Run()
  