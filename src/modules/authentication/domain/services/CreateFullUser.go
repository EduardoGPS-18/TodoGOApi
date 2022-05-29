package services

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
)

func CreateFullUser(email, password string) (*entities.User, errors.DomainError) {
	userCredentials, err := entities.NewUserCredentials(email, password)
	if err != nil {
		return nil, err
	}

	userAuthorization, err := entities.NewUserAuthorization(userCredentials.ID)
	if err != nil {
		return nil, err
	}

	user := entities.NewUser(
		userCredentials,
		userAuthorization,
	)

	return user, nil
}
