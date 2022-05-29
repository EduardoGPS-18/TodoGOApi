package entities

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Value []byte
}

func NewPassword(value string) (*Password, errors.DomainError) {
	if len(value) < 6 {
		return nil, errors.NewValidationError("Password weak")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(value), 10)

	password := Password{
		Value: hashedPassword,
	}

	return &password, nil
}

func (p *Password) ComparesPassword(password string) errors.DomainError {
	err := bcrypt.CompareHashAndPassword(p.Value, []byte(password))
	if err != nil {
		return errors.NewValidationError("Password not matches")
	}
	return nil
}
