package orm

type UserCredentialsORM struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
