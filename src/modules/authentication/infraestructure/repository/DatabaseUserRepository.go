package repository

import (
	"errors"

	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/database/mapper"
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/database/orm"
)

type DatabaseUserRepository struct {
	Database          Database
	userAuthorization []entities.UserAuthorization
}

func (r *DatabaseUserRepository) SaveUserCredentials(user *entities.UserCredentials) error {
	credentialsUser := mapper.UserCredentialsToORM(*user)
	r.Database.Save(credentialsUser)
	return nil
}

func (r *DatabaseUserRepository) FindUserCredentials(email string) (*entities.UserCredentials, error) {
	var user orm.UserCredentialsORM
	r.Database.FindOne(&user, "email = ?", email)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}
	return mapper.UserCredentialsFromORM(user), nil
}

func (d *DatabaseUserRepository) SaveUserAuthorization(userAuthorization *entities.UserAuthorization) error {
	for i, userAuth := range d.userAuthorization {
		if userAuth.UserID == userAuthorization.UserID {
			d.userAuthorization[i] = *userAuthorization
		}
	}
	d.userAuthorization = append(d.userAuthorization, *userAuthorization)
	return nil
}

func (d *DatabaseUserRepository) FindUserAuthorization(userID string) (*entities.UserAuthorization, error) {
	for _, user := range d.userAuthorization {
		if user.UserID == userID {
			return &user, nil
		}
	}
	return nil, errors.New("not founded user")
}

func NewDatabaseUserRepository(database Database) entities.UserRepository {
	return &DatabaseUserRepository{
		Database: database,
	}
}
