package database

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/infraestructure/database/orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *PostgresDBAdapter

type PostgresDBAdapter struct {
	PostgresConnection *gorm.DB
}

// TODO: Fix this file struct

func InitializePG() error {
	psqlConnectionString := "postgresql://postgres:postgres_secret@localhost/tododb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(psqlConnectionString))
	if err != nil {
		panic("error on connect on database")
	}

	db.Table("user_credentials").AutoMigrate(orm.UserCredentialsORM{})

	Instance = &PostgresDBAdapter{
		PostgresConnection: db.Table("user_credentials"),
	}
	return nil
}

func (p *PostgresDBAdapter) Save(data interface{}) {
	psqlConnectionString := "postgresql://postgres:postgres_secret@localhost/tododb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(psqlConnectionString))
	if err != nil {
		panic("error on connect on database")
	}

	db.Table("user_credentials").Create(data)
}

func (p *PostgresDBAdapter) FindOne(assignVariable interface{}, condition, value string) {
	psqlConnectionString := "postgresql://postgres:postgres_secret@localhost/tododb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(psqlConnectionString))
	if err != nil {
		panic("error on connect on database")
	}
	db.Table("user_credentials").Where(condition, value).First(&assignVariable)
}

func (p *PostgresDBAdapter) FindAll(assignVariable interface{}) {
	p.PostgresConnection.Find(assignVariable)
}

func (p *PostgresDBAdapter) UpdateWhere(id string, data interface{}) {
	p.PostgresConnection.Find("id = ?", id)
	p.PostgresConnection.Save(data)
}

func (p *PostgresDBAdapter) DeleteWhere(condition, value string) {
	p.PostgresConnection.Delete(condition, value)
}
