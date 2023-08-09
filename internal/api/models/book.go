package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `gorm:"not null"`
	ISBN   string `gorm:"not null;primaryKey"`
	Price  float64
	UserId int
}
