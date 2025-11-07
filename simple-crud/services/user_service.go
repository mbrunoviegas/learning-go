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

func (s *UserService) CreateUser(user models.User) models.ID {
	return s.db.AddUser(user)
}
