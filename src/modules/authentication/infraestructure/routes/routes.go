package routes

import (
	"log"

	"com.task-go-api.com/dudu.com/src/modules/authentication/application/services"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	database "com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/database/connections"
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/helpers"
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/repository"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/controller"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/dtos"
	"github.com/gin-gonic/gin"
)

func DefineAuthenticationRoutes(r *gin.Engine) *gin.Engine {
	err := database.InitializePG()
	if err != nil {
		log.Fatal(err)
	}

	var userRepository entities.UserRepository = repository.NewDatabaseUserRepository(database.Instance)
	var loginService services.LoginService = services.NewLoginService(userRepository)
	var registerService services.RegisterService = services.NewRegisterService(userRepository)

	r.Group("/auth")
	{
		r.POST("/login", func(ctx *gin.Context) {
			helpers.GinRouteAdapter(ctx, func(body controller.LoginUserDTO) (*dtos.UserDTO, *dtos.ErrorDto) {
				return controller.NewLoginController(loginService).Handle(body)
			})
		})

		r.POST("/register", func(ctx *gin.Context) {
			helpers.GinRouteAdapter(ctx, func(body controller.RegisterUserDTO) (*dtos.UserDTO, *dtos.ErrorDto) {
				return controller.NewRegisterController(registerService).Handle(body)
			})
		})
	}

	return r
}
