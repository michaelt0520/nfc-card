package models

import "gorm.io/gorm"

// Information struct
type Information struct {
	gorm.Model
	UserID   uint   `gorm:"column:user_id" json:"user_id"`
	Name     string `gorm:"column:name" json:"name"`
	Data     string `gorm:"column:data" json:"data"`
	Visibled bool   `gorm:"column:visibled;default:true" json:"visibled"`
	User     User   `json:"user"`
}
