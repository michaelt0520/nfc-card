package models

import "gorm.io/gorm"

// Company struct
type Company struct {
	gorm.Model
	Name    string `gorm:"column:name" json:"name"`
	Address string `gorm:"column:address" json:"address"`
	Logo    string `gorm:"column:logo" json:"logo"`
	Website string `gorm:"column:website" json:"website"`
}
