package main

import (
	"context"

	"github.com/ikiselewskii/avito-test-task/database"
	"github.com/ikiselewskii/avito-test-task/utils"
	"github.com/ikiselewskii/avito-test-task/webserver"
)

func main() {
	router := webserver.CreateRouterEngine()
	webserver.InitializeEndpoints(router)
	database.InitializeDBConnection(utils.SerializeDSN())
	database.CreateTables(context.Background())
	router.Run()
}
