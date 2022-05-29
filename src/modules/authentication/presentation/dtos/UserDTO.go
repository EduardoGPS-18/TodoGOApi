package dtos

import "com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"

type UserDTO struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

func UserDTOFromEntity(user *entities.User) *UserDTO {
	return &UserDTO{
		ID:          user.UserCredentials.ID,
		Email:       user.UserCredentials.Email.Value,
		AccessToken: user.UserAuthorization.JWT.Value,
	}
}
