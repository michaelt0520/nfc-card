package models

import "gorm.io/gorm"

// Card struct
type Card struct {
	gorm.Model
	Code      string `gorm:"column:code;unique" json:"code"`
	UserID    uint   `gorm:"column:user_id" json:"user_id"`
	Activated bool   `gorm:"column:activated;default:true" json:"activated"`
	User      User   `json:"user"`
}
