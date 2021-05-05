package repositories

import (
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

// var _ Repository = &UserRepository{}

// Find : get card by code
func (u *UserRepository) All(result interface{}) error {
	query := u.UserTable().Preload("Company").Find(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Find : get user
func (u *UserRepository) Find(result interface{}, data map[string]interface{}) error {
	query := u.UserTable().Preload("Company").Preload("Informations").Where(data).First(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Where :
func (u *UserRepository) Where(result interface{}, data map[string]interface{}) error {
	query := u.UserTable().Preload("Company").Preload("Informations").Where(data).Find(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Create : Save user to db
func (u *UserRepository) Create(model interface{}) error {
	record := model.(*models.User)
	if err := record.HashPassword(record.Password); err != nil {
		return err
	}

	query := u.UserTable().Where("email = ?", record.Email).Or("username = ?", record.Username).FirstOrCreate(record)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Update : Update user to db
func (u *UserRepository) Update(model interface{}, data map[string]interface{}) error {
	record := model.(*models.User)

	var query *gorm.DB
	if password, ok := data["password"]; ok {
		if err := record.HashPassword(password.(string)); err != nil {
			return err
		}

		query = u.UserTable().Save(record)
		if err := query.Error; err != nil {
			return err
		}
	} else {
		query = u.UserTable().Model(record).Updates(data)
		if err := query.Error; err != nil {
			return err
		}
	}

	return nil
}

// Destroy : destroy category
func (u *UserRepository) Destroy(model interface{}) error {
	record := model.(*models.User)

	query := u.UserTable().Delete(record)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}
