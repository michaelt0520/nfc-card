package models

import "time"

// User struct
type User struct {
	gorm.Model
	Email	 	 			string					`gorm:"column:email"       json:"email"`
	Password	  	string					`gorm:"column:password"  	 json:"password"`
	Username 			string					`gorm:"column:username"	 	 json:"username"`
	Name  				string					`gorm:"column:name"				 json:"name"`
	AvatarData    string	  			`gorm:"column:avatar_data" json:"avatar_data"`
	AuthProviders []*AuthProvider `gorm:"foreignKey:UserID"  json:"auth_providers"`
}
