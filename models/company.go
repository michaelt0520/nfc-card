package models

import "gorm.io/gorm"

// Company struct
type Company struct {
	gorm.Model
	Name    string  `gorm:"column:name;unique;not null" json:"name"`
	Address string  `gorm:"column:address" json:"address"`
	Logo    string  `gorm:"column:logo" json:"logo"`
	Website string  `gorm:"column:website" json:"website"`
	Hotline string  `gorm:"column:hotline" json:"hotline"`
	Users   []*User `gorm:"foreignKey:CompanyID" json:"users"`
	Cards   []*Card `gorm:"foreignKey:CompanyID" json:"cards"`
}
