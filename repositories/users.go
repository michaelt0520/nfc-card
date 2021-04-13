package repositories

import "github.com/michaelt0520/nfc-card/models"

// UserRepository : struct
type UserRepository struct{}

// NewUserRepository ...
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Find : get user by username
func (u *UserRepository) Find(data map[string]interface{}) (*models.User, error) {
	var user models.User

	if username, ok := data["username"]; ok {
		if err := GetDB().Preload("Informations").Where("username = ?", username).First(&user).Error; err != nil {
			return nil, err
		}

		return &user, nil
	} else {
		if err := GetDB().Preload("Informations").Where(data).First(&user).Error; err != nil {
			return nil, err
		}

		return &user, nil
	}
}

// Create : Save user to db
func (repo *UserRepository) Create(user *models.User) error {
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}

	result := GetDB().Where("email = ?", user.Email).Or("username = ?", user.Username).FirstOrCreate(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
