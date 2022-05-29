package controller

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/application/services"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/dtos"
)

type RegisterUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"`
}

type RegisterController struct {
	registerService services.RegisterService
}

func (r *RegisterController) Handle(createUserDTO RegisterUserDTO) (*entities.User, *dtos.ErrorDto) {
	user, err := r.registerService.RegisterNewUser(createUserDTO.Email, createUserDTO.Password)
	if err != nil {
		return nil, dtos.NewErrorDto(err.Message())
	}

	return user, nil
}

func NewRegisterController(registerService services.RegisterService) *RegisterController {
	return &RegisterController{
		registerService: registerService,
	}
}
