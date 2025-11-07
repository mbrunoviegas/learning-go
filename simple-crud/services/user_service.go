package services

import (
	"simple-crud/models"
	"simple-crud/storage"
)

type UserService struct {
	db *storage.UserDb
}

func NewUserService(db *storage.UserDb) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user models.User) (models.ID, error) {
	return s.db.AddUser(user)
}

func (s *UserService) GetUserById(id models.ID) (models.User, bool) {
	return s.db.GetUser(id)
}

func (s *UserService) ListUsers() []models.User {
	return s.db.ListUsers()
}

func (s *UserService) UpdateUser(id models.ID, updateData models.User) (models.User, bool) {
	return s.db.UpdateUser(id, updateData)
}

func (s *UserService) DeleteUser(id models.ID) (models.User, bool) {
	return s.db.DeleteUser(id)
}
