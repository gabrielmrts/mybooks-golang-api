package models

type EmailVerification struct {
	Email string `json:"-" gorm:"primarykey;index"`
	Code  string `json:"code" gorm:"index"`
}
