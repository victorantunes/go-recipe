package services

import (
	"errors"
	"go-recipe/cmd/pkg/validator"
	"go-recipe/db/repositories"
	"go-recipe/models"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Validate struct fields
	if err := validator.Validate.Struct(user); err != nil {
		return err
	}

	// Check uniqueness of email (business logic)
	existingUser, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email already in use")
	}

	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) ListUsers(companyId string) ([]models.User, error) {
	return s.repo.List(companyId)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
