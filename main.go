package main

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.DefineAuthenticationRoutes(r)
	r.Run()

}
