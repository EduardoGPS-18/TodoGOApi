package repository

import (
	"errors"

	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
)

type LocalUserRepository struct {
	usersCredentials  []entities.UserCredentials
	userAuthorization []entities.UserAuthorization
}

func (r *LocalUserRepository) SaveUserCredentials(user *entities.UserCredentials) error {
	r.usersCredentials = append(r.usersCredentials, *user)
	return nil
}

func (r *LocalUserRepository) FindUserCredentials(email string) (*entities.UserCredentials, error) {
	for _, user := range r.usersCredentials {
		if user.Email.Value == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *LocalUserRepository) SaveUserAuthorization(userAuthorization *entities.UserAuthorization) error {
	for i, userAuth := range r.userAuthorization {
		if userAuth.UserID == userAuthorization.UserID {
			r.userAuthorization[i] = *userAuthorization
		}
	}
	r.userAuthorization = append(r.userAuthorization, *userAuthorization)
	return nil
}

func (r *LocalUserRepository) FindUserAuthorization(userID string) (*entities.UserAuthorization, error) {
	for _, user := range r.userAuthorization {
		if user.UserID == userID {
			return &user, nil
		}
	}
	return nil, errors.New("not founded user")
}

func NewLocalUserRepository() entities.UserRepository {
	return &LocalUserRepository{}
}
