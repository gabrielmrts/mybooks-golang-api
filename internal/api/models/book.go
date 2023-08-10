package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string  `json:"title" gorm:"not null"`
	ISBN   string  `json:"isbn" gorm:"not null;primaryKey"`
	Price  float64 `json:"price" gorm:"default:0"`
	UserId int
}
