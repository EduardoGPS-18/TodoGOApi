package repository

type Database interface {
	Save(data interface{})

	FindOne(assignVariable interface{}, condition, value string)
	FindAll(assignVariable interface{})

	UpdateWhere(id string, data interface{})

	DeleteWhere(condition, value string)
}
