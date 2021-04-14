package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CardType : card type
type CardType uint

const (
	Personal CardType = iota + 1
	Business
)

// UserRole : user role
type UserRole uint

const (
	UserMember UserRole = iota + 1
	UserAdmin
)

// User struct
type User struct {
	gorm.Model
	Name         string         `gorm:"column:name" json:"name"`
	Avatar       string         `gorm:"column:avatar" json:"avatar"`
	Username     string         `gorm:"column:username;unique;not null" json:"username"`
	Email        string         `gorm:"column:email;unique;not null" json:"email"`
	Password     string         `gorm:"column:password" json:"password"`
	Type         CardType       `gorm:"column:type" json:"type"`
	Role         UserRole       `gorm:"column:role" json:"role"`
	JWT          string         `gorm:"column:jwt" json:"jwt"`
	CompanyID    uint           `gorm:"column:company_id" json:"company_id"`
	Cards        []*Card        `gorm:"foreignKey:UserID" json:"cards"`
	Informations []*Information `gorm:"foreignKey:UserID" json:"informations"`
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

// IsMember check role user is member
func (user *User) IsMember() bool {
  return user.Role == UserMember
}

// IsAdmin check role user is admin
func (user *User) IsAdmin() bool {
  return user.Role == UserAdmin
}
