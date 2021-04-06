package models

import "gorm.io/gorm"

// SocialInformation struct
type SocialInformation struct {
	gorm.Model
	UserID    uint32 `gorm:"column:user_id"    json:"user_id"`
	Provider  string `gorm:"column:provider"   json:"provider"`
	Data      string `gorm:"column:data"	     json:"data"`
	IsPrivate bool   `gorm:"column:is_private" json:"is_private"`
}
