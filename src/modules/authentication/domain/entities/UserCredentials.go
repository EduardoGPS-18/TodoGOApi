package entities

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
	"github.com/google/uuid"
)

type UserCredentials struct {
	ID       string
	Email    Email
	Password Password
}

func NewUserCredentials(email, password string) (*UserCredentials, errors.DomainError) {

	userEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	userPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	user := UserCredentials{
		ID:       uuid.NewString(),
		Email:    *userEmail,
		Password: *userPassword,
	}

	return &user, nil
}
