package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string  `json:"title" gorm:"not null"`
	ISBN   string  `json:"isbn" gorm:"unique"`
	Price  float64 `json:"price"`
	UserId int
}
