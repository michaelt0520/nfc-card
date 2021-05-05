package repositories

import (
	"github.com/michaelt0520/nfc-card/models"
	"github.com/mitchellh/mapstructure"
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

var _ Repository = &UserRepository{}

// Find : get card by code
func (u *UserRepository) All(result interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.UserTable().Scopes(scopes...).Preload("Company").Find(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Find : get user
func (u *UserRepository) Find(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.UserTable().Scopes(scopes...).Preload("Company").Preload("Informations").Where(data).First(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Where :
func (u *UserRepository) Where(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	var fields models.User
	mapstructure.Decode(data, &fields)

	query := u.UserTable().Scopes(scopes...).Preload("Company").Preload("Informations").Where(fields, "company_id")
	if err := query.Error; err != nil {
		return nil, err
	}

	return query.Find(result), nil
}

func (u *UserRepository) Search(db *gorm.DB, keyword string) (*gorm.DB, error) {
	db.Where("lower(email) LIKE ?", keyword)
	db.Or("lower(name) LIKE ?", keyword)
	db.Or("lower(phone_number) LIKE ?", keyword)

	return db, nil
}

// Create : Save user to db
func (u *UserRepository) Create(model interface{}) (*gorm.DB, error) {
	record := model.(*models.User)
	if err := record.HashPassword(record.Password); err != nil {
		return nil, err
	}

	query := u.UserTable().Where("email = ?", record.Email).Or("username = ?", record.Username).FirstOrCreate(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Update : Update user to db
func (u *UserRepository) Update(model interface{}, data map[string]interface{}) (*gorm.DB, error) {
	record := model.(*models.User)

	var query *gorm.DB
	if password, ok := data["password"]; ok {
		if err := record.HashPassword(password.(string)); err != nil {
			return nil, err
		}

		query = u.UserTable().Save(record)
		if err := query.Error; err != nil {
			return nil, err
		}
	} else {
		query = u.UserTable().Model(record).Updates(data)
		if err := query.Error; err != nil {
			return nil, err
		}
	}

	return query, nil
}

// Destroy : destroy category
func (u *UserRepository) Destroy(model interface{}) (*gorm.DB, error) {
	record := model.(*models.User)

	query := u.UserTable().Delete(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}
