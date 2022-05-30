package controller

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/application/services"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/dtos"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/helpers"
)

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginController struct {
	loginService services.LoginService
}

func (l *LoginController) Handle(LoginUserDTO LoginUserDTO) (*dtos.UserDTO, *dtos.ErrorDto) {
	user, err := l.loginService.LoginWithEmailAndPassword(LoginUserDTO.Email, LoginUserDTO.Password)
	if err != nil {
		return nil, helpers.ConvertDomainErrorToDTO(err)
	}
	return dtos.UserDTOFromEntity(user), nil
}

func NewLoginController(loginService services.LoginService) *LoginController {
	return &LoginController{
		loginService: loginService,
	}
}
