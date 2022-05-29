package entities

import (
	"errors"
	"time"

	domainErrors "com.task-go-api.com/dudu.com/src/modules/authentication/domain/errors"
	"github.com/golang-jwt/jwt/v4"
)

type JWTPaylod struct {
	jwt.RegisteredClaims
	ID string
}

type JWT struct {
	Value string
}

func NewJWT(id string) (*JWT, domainErrors.DomainError) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 5)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        id,
		Issuer:    "TODO-GO-APi",
	}
	payload := JWTPaylod{
		ID:               id,
		RegisteredClaims: claims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)
	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return nil, domainErrors.NewValidationError("error on sign token")
	}

	jwt := JWT{
		Value: tokenString,
	}
	return &jwt, nil
}

func VerifyJWTValidation(token string) (*JWTPaylod, domainErrors.DomainError) {
	sessionExpiredError := domainErrors.NewSessionExpiredError("Expired session!")

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("your session has expired")
		}
		return []byte("secretKey"), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JWTPaylod{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, jwt.ErrTokenExpired) {
			return nil, sessionExpiredError
		}
		return nil, sessionExpiredError
	}

	payload, ok := jwtToken.Claims.(*JWTPaylod)
	if !ok {
		return nil, sessionExpiredError
	}

	return payload, nil
}
