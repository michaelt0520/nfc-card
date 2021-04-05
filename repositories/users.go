package repositories

import "github.com/michaelt0520/nfc-card/models"

// UserRepository : struct
type UserRepository struct{}

// NewUserRepository ...
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Find : get user by username
func (u *UserRepository) Find(username string) (*models.User, error) {
	var user models.User

	if err := GetDB().Where("username = ?", username).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
