package models

type User struct {
	ID uint `json:"id" gorm:"primarykey"`

	Role    string  `json:"role" gorm:"default:USER"`
	Account Account `gorm:"ForeignKey:UserId"`
	Books   []Book  `gorm:"ForeignKey:UserId"`
}
