package services

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
)

type RegisterService struct {
	userRepository entities.UserRepository
}

func (r *RegisterService) RegisterNewUser(email, password string) (*entities.User, errors.DomainError) {
	userWithSameEmail, _ := r.userRepository.FindUserCredentials(email)
	if userWithSameEmail != nil {
		return nil, errors.NewValidationError("user already exists")
	}

	userCredentials, err := entities.NewUserCredentials(email, password)
	if err != nil {
		return nil, err
	}

	repositoryError := r.userRepository.SaveUserCredentials(userCredentials)
	if repositoryError != nil {
		return nil, err
	}

	userAuthorization, err := entities.NewUserAuthorization(userCredentials.ID)
	if err != nil {
		return nil, err
	}

	repositoryError = r.userRepository.SaveUserAuthorization(userAuthorization)
	if repositoryError != nil {
		return nil, err
	}

	user := entities.NewUser(userCredentials, userAuthorization)
	return user, nil
}

func NewRegisterService(
	userRepository entities.UserRepository,
) RegisterService {
	return RegisterService{
		userRepository: userRepository,
	}
}
