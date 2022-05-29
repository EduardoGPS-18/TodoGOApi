package routes

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/application/services"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/helpers"
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/repository"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/controller"
	"github.com/gin-gonic/gin"
)

func ConvertErrorToStatusCode(err interface{}) int {
	switch err.(type) {
	case errors.ValidationError:
		return 400
	case errors.SessionExpiredError:
		return 403
	default:
		return 500
	}
}

func DefineAuthenticationRoutes(r *gin.Engine) *gin.Engine {
	var userRepository entities.UserRepository = repository.NewLocalUserRepository()
	var loginService services.LoginService = services.NewLoginService(userRepository)

	var registerService services.RegisterService = services.NewRegisterService(userRepository)

	r.Group("/auth")
	{
		// TODO: TRY TO MOVE DUPLICATED CODE TO A HELPER METHOD
		r.POST("/login", func(ctx *gin.Context) {
			loginUserDTO := helpers.GetGinBody[controller.LoginUserDTO](ctx)
			user, err := controller.NewLoginController(loginService).Handle(loginUserDTO)

			if err != nil {
				statusCode := ConvertErrorToStatusCode(err)
				ctx.JSON(statusCode, err)
			} else {
				ctx.JSON(200, user.ToJson())
			}
		})
		r.POST("/register", func(ctx *gin.Context) {
			registerUserDTO := helpers.GetGinBody[controller.RegisterUserDTO](ctx)
			user, err := controller.NewRegisterController(registerService).Handle(registerUserDTO)
			if err != nil {
				statusCode := ConvertErrorToStatusCode(err)
				ctx.JSON(statusCode, err)
			} else {
				ctx.JSON(200, user.ToJson())
			}
		})
	}

	return r
}
