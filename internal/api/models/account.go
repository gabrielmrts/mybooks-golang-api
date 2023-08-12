package models

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        uint         `json:"account_id" gorm:"primarykey"`
	Email     string       `json:"email" gorm:"not null"`
	Password  string       `json:"-" gorm:"not null"`
	UserId    uint         `json:"-"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}
