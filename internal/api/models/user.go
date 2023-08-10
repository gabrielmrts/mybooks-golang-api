package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"-" gorm:"not null"`
	Role     string `json:"role" gorm:"default:USER"`
	Books    []Book `gorm:"ForeignKey:UserId"`
}
