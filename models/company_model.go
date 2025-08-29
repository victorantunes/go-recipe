package models

type Company struct {
	BaseModel
	Name  string `json:"name" validate:"required"`
	Users []User `json:"users" gorm:"foreignKey:CompanyID"`
}
