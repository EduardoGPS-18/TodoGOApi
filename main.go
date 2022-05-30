package main

import (
	database "com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/database/connections"
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitializePG()

	routes.DefineAuthenticationRoutes(r)

	r.Run()
}
