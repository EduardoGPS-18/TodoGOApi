package services

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
)

type LoginService struct {
	userRepository entities.UserRepository
}

func (l *LoginService) LoginWithEmailAndPassword(email, password string) (*entities.User, errors.DomainError) {
	userCredentials, _ := l.userRepository.FindUserCredentials(email)
	if userCredentials == nil {
		return nil, errors.NewValidationError("invalid email or password")
	}

	err := userCredentials.Password.ComparesPassword(password)
	if err != nil {
		return nil, errors.NewValidationError("invalid email or password")
	}

	userAuthorization, err := entities.NewUserAuthorization(userCredentials.ID)
	if err != nil {
		return nil, err
	}

	repositoryError := l.userRepository.SaveUserAuthorization(userAuthorization)
	if repositoryError != nil {
		return nil, err
	}

	return entities.NewUser(userCredentials, userAuthorization), nil
}

func NewLoginService(
	userCredentialsRepository entities.UserRepository,
) LoginService {

	return LoginService{
		userRepository: userCredentialsRepository,
	}
}
