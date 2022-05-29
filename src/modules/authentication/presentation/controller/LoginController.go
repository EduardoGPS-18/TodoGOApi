package controller

import (
	"reflect"

	"com.task-go-api.com/dudu.com/src/modules/authentication/application/services"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/dtos"
)

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginController struct {
	loginService services.LoginService
}

func (l *LoginController) Handle(LoginUserDTO LoginUserDTO) (*entities.User, *dtos.ErrorDto) {
	user, err := l.loginService.LoginWithEmailAndPassword(LoginUserDTO.Email, LoginUserDTO.Password)
	if err != nil {
		return nil, dtos.NewErrorDto(err.Message())
	}
	reflect.TypeOf(err)
	return user, nil
}

func NewLoginController(loginService services.LoginService) *LoginController {
	return &LoginController{
		loginService: loginService,
	}
}
