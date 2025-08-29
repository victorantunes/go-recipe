package dto

import "go-recipe/models"

type CompanyResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Users []user `json:"users" gorm:"foreignKey:CompanyID"`
}

type user struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func MapCompaniesToResponse(company []models.Company) []CompanyResponse {
	var usersDTO = make([]CompanyResponse, 0, len(company))

	for _, u := range company {
		usersDTO = append(usersDTO, MapCompanyToResponse(u))
	}

	return usersDTO
}

func MapCompanyToResponse(company models.Company) CompanyResponse {
	var users = make([]user, 0, len(company.Users))

	for _, u := range company.Users {
		users = append(users, user{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return CompanyResponse{
		ID:    company.ID,
		Name:  company.Name,
		Users: users,
	}
}
