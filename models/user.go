package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" gorm:"uniqueIndex" validate:"required,email"`
}
