package services

import (
	"cql-backend/db/repositories"
	"cql-backend/models"
	"errors"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Here you can add business logic like validation or checks
	if user.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.List()
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
