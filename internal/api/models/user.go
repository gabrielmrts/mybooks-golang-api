package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
	Name      string       `json:"name" gorm:"not null"`
	Email     string       `json:"email" gorm:"not null"`
	Password  string       `json:"-" gorm:"not null"`
	Role      string       `json:"role" gorm:"default:USER"`
	Books     []Book       `gorm:"ForeignKey:UserId"`
}
