package database

import (
	"gorm.io/gorm"
)

type MongoDB struct {
	DB gorm.DB
}

//TODO: Implement Mongo database
func (m *MongoDB) InitializeDB() {

}
