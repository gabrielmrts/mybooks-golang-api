package models

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        uint         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
	Title     string       `json:"title" gorm:"not null"`
	ISBN      string       `json:"isbn" gorm:"unique"`
	Price     float64      `json:"price" gorm:"default:0"`
	Quantity  uint         `json:"quantity" gorm:"default:0"`
	UserId    uint         `json:"-"`
}
