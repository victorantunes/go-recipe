package dto

import "go-recipe/models"

type UserResponse struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Company company `json:"company"`
}

type company struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func MapUsersToResponse(users []models.User) []UserResponse {
	var usersDTO = make([]UserResponse, 0, len(users))

	for _, u := range users {
		usersDTO = append(usersDTO, MapUserToResponse(u))
	}

	return usersDTO
}

func MapUserToResponse(user models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Company: company{
			ID:   user.Company.ID,
			Name: user.Company.Name,
		},
	}
}
