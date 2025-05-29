package db

import (
	"cql-backend/db/repositories"
	"gorm.io/gorm"
)

type Store struct {
	UserRepository *repositories.UserRepository
	// Add other repos here in the future
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		UserRepository: repositories.NewUserRepository(db),
	}
}
