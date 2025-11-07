package storage

import (
	"maps"
	"simple-crud/models"
	"slices"

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

func (db *UserDb) ListUsers() []models.User {
	users := slices.Collect(maps.Values(db.data))

	return users
}

func (db *UserDb) UpdateUser(id models.ID, updateData models.User) (models.User, bool) {
	user, exists := db.data[id]

	if !exists {
		return user, exists
	}

	db.data[id] = updateData

	return db.data[id], exists
}

func (db *UserDb) DeleteUser(id models.ID) (models.User, bool) {
	user, exists := db.data[id]

	delete(db.data, id)

	return user, exists
}
