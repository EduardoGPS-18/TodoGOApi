package controller

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/application/services"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/dtos"
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/helpers"
)

type RegisterUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"`
}

type RegisterController struct {
	registerService services.RegisterService
}

func (r *RegisterController) Handle(createUserDTO RegisterUserDTO) (*dtos.UserDTO, *dtos.ErrorDto) {
	user, err := r.registerService.RegisterNewUser(createUserDTO.Email, createUserDTO.Password)
	if err != nil {
		return nil, helpers.ConvertDomainErrorToDTO(err)
	}
	return dtos.UserDTOFromEntity(user), nil
}

func NewRegisterController(registerService services.RegisterService) *RegisterController {
	return &RegisterController{
		registerService: registerService,
	}
}
