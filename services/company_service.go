package services

import (
	"go-recipe/cmd/pkg/validator"
	"go-recipe/db/repositories"
	"go-recipe/models"
)

type CompanyService struct {
	repo *repositories.CompanyRepository
}

func NewCompanyService(repo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) CreateCompany(company *models.Company) error {
	if err := validator.Validate.Struct(company); err != nil {
		return err
	}

	return s.repo.Create(company)
}

func (s *CompanyService) GetCompanyByID(id uint) (*models.Company, error) {
	return s.repo.GetByID(id)
}

func (s *CompanyService) ListCompanies() ([]models.Company, error) {
	return s.repo.List()
}

func (s *CompanyService) DeleteCompany(id uint) error {
	return s.repo.Delete(id)
}
