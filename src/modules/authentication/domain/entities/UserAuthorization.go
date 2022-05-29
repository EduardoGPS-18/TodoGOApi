package entities

import "com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"

type UserAuthorization struct {
	UserID string
	JWT    JWT
}

func NewUserAuthorization(userID string) (*UserAuthorization, errors.DomainError) {
	token, err := NewJWT(userID)
	if err != nil {
		return nil, err
	}

	userAuth := UserAuthorization{
		UserID: userID,
		JWT:    *token,
	}

	return &userAuth, nil
}
