package entities

import (
	"regexp"

	"com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
)

type Email struct {
	Value string
}

func NewEmail(value string) (*Email, errors.DomainError) {
	emailRegexp, _ := regexp.Compile(`^[a-z]+(\w+)@([a-z]+)\.([a-z]+){1,3}`)
	if !emailRegexp.MatchString(value) {
		return nil, errors.NewValidationError("UserCredentials email must be valid")
	}

	email := &Email{
		Value: value,
	}
	return email, nil
}
