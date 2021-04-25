package models

import "gorm.io/gorm"

// Contact struct
type Contact struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Email       string `gorm:"column:email" json:"email"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
}
