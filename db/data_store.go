package db

import (
	"go-recipe/db/repositories"

	"gorm.io/gorm"
)

type DataStore struct {
	UserRepository    *repositories.UserRepository
	CompanyRepository *repositories.CompanyRepository
	// Add other repos here in the future
}

func NewDataStore(db *gorm.DB) *DataStore {
	return &DataStore{
		UserRepository:    repositories.NewUserRepository(db),
		CompanyRepository: repositories.NewCompanyRepository(db),
	}
}
