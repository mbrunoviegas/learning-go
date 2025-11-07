package storage

import (
	"simple-crud/models"

	"github.com/google/uuid"
)

type UserDb struct {
	data map[models.ID]models.User
}

func NewUserDb() *UserDb {
	return &UserDb{
		data: make(map[models.ID]models.User),
	}
}

func (db *UserDb) AddUser(user models.User) models.ID {
	id := models.ID(uuid.New().String())
	db.data[id] = user

	return id
}

func (db *UserDb) GetUser(id models.ID) (models.User, bool) {
	user, exists := db.data[id]

	return user, exists
}
