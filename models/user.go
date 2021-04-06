package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Email              string               `gorm:"column:email"       json:"email"`
	Password           string               `gorm:"column:password"  	 json:"password"`
	Username           string               `gorm:"column:username"	 	 json:"username"`
	Name               string               `gorm:"column:name"				 json:"name"`
	AvatarData         string               `gorm:"column:avatar_data" json:"avatar_data"`
	SocialInformations []*SocialInformation `gorm:"foreignKey:UserID"  json:"social_informations"`
}

// HashPassword : encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return err
	}
	user.Password = string(bytes)

	return nil
}

// CheckPassword : checks user password
func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
