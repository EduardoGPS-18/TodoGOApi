package entities

type UserRepository interface {
	SaveUserCredentials(user *UserCredentials) error
	FindUserCredentials(email string) (*UserCredentials, error)

	SaveUserAuthorization(userAuthorization *UserAuthorization) error
	FindUserAuthorization(userID string) (*UserAuthorization, error)
}

type User struct {
	UserCredentials   *UserCredentials
	UserAuthorization *UserAuthorization
}

func NewUser(userCredentials *UserCredentials, userAuthorization *UserAuthorization) *User {
	return &User{
		UserAuthorization: userAuthorization,
		UserCredentials:   userCredentials,
	}
}

func (u *User) ToJson() map[string]string {
	return map[string]string{
		"id":          u.UserAuthorization.UserID,
		"email":       u.UserCredentials.Email.Value,
		"accessToken": u.UserAuthorization.JWT.Value,
	}
}
