package repositories

import (
	"go-recipe/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) GetByID(id uint) (*models.Company, error) {
	var company models.Company
	if err := r.db.Preload("Users").First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepository) List() ([]models.Company, error) {
	var companies []models.Company
	err := r.db.Preload("Users").Find(&companies).Error
	return companies, err
}

func (r *CompanyRepository) Delete(id uint) error {
	var company models.Company
	if err := r.db.First(&company, id).Error; err != nil {
		return err
	}
	return r.db.Unscoped().Delete(&models.Company{}, &id).Error
}
