package models

type User struct {
	BaseModel
	Name      string  `json:"name" validate:"required"`
	Email     string  `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	CompanyID int     `json:"company_id" validate:"required"`
	Company   Company `json:"company" validate:"-"`
}
