package services

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/services"
)

type RegisterService struct {
	userRepository entities.UserRepository
}

func (r *RegisterService) RegisterNewUser(email, password string) (*entities.User, errors.DomainError) {
	userWithSameEmail, _ := r.userRepository.FindUserCredentials(email)
	if userWithSameEmail != nil {
		return nil, errors.NewValidationError("user already exists")
	}

	user, err := services.CreateFullUser(email, password)
	if err != nil {
		return nil, err
	}
	repositoryError := r.userRepository.SaveUserCredentials(user.UserCredentials)
	if repositoryError != nil {
		return nil, err
	}

	repositoryError = r.userRepository.SaveUserAuthorization(user.UserAuthorization)
	if repositoryError != nil {
		return nil, err
	}

	return user, nil
}

func NewRegisterService(
	userRepository entities.UserRepository,
) RegisterService {
	return RegisterService{
		userRepository: userRepository,
	}
}
