package mapper

import (
	"fmt"

	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/entities"
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/database/orm"
)

func UserCredentialsToORM(user entities.UserCredentials) *orm.UserCredentialsORM {
	fmt.Printf("SENHA ARMAZENADA %v\n\n", string(user.Password.Value))
	return &orm.UserCredentialsORM{
		ID:       user.ID,
		Email:    user.Email.Value,
		Password: string(user.Password.Value),
	}
}

func UserCredentialsFromORM(user orm.UserCredentialsORM) *entities.UserCredentials {
	password := entities.Password{Value: []byte(user.Password)}
	email := entities.Email{Value: user.Email}

	return &entities.UserCredentials{
		ID:       user.ID,
		Email:    email,
		Password: password,
	}
}
