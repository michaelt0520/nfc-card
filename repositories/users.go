package repositories

import (
	"github.com/michaelt0520/nfc-card/interfaces"
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

// UserRepository : struct
type UserRepository struct{}

// NewUserRepository ...
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) UserTable() *gorm.DB {
	return GetDB().Model(&models.User{})
}

var _ interfaces.IUserRepository = &UserRepository{}

// Find : get user
func (u *UserRepository) Find(result *models.User, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.UserTable().Scopes(scopes...).First(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Where :
func (u *UserRepository) Where(result *[]models.User, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.UserTable().Scopes(scopes...).Find(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Create : Save user to db
func (u *UserRepository) Create(model *models.User) error {
	if err := model.HashPassword(model.Password); err != nil {
		return err
	}

	query := u.UserTable().Where("email = ?", model.Email).Or("username = ?", model.Username).FirstOrCreate(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Update : Update user to db
func (u *UserRepository) Update(model *models.User, data map[string]interface{}) error {
	var query *gorm.DB
	if password, ok := data["password"]; ok {
		if err := model.HashPassword(password.(string)); err != nil {
			return err
		}

		query = u.UserTable().Save(model)
		if err := query.Error; err != nil {
			return err
		}
	} else {
		query = u.UserTable().Model(model).Updates(data)
		if err := query.Error; err != nil {
			return err
		}
	}

	return nil
}

// Destroy : destroy category
func (u *UserRepository) Destroy(model *models.User) error {
	query := u.UserTable().Delete(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}
