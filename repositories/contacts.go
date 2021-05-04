package repositories

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

// ContactRepository : struct
type ContactRepository struct{}

// NewContactRepository ...
func NewContactRepository() *ContactRepository {
	return &ContactRepository{}
}

func (u *ContactRepository) ContactTable() *gorm.DB {
	return GetDB().Model(&models.Contact{})
}

var _ Repository = &ContactRepository{}

// Find : get card by code
func (u *ContactRepository) All(result interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.ContactTable().Scopes(scopes...).Find(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Find : get contact by id
func (u *ContactRepository) Find(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.ContactTable().Scopes(scopes...).Where(data).First(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Create : create contact
func (u *ContactRepository) Where(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.ContactTable().Scopes(scopes...).Where(data).Find(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Create : Save user to db
func (u *ContactRepository) Create(model interface{}) (*gorm.DB, error) {
	record := model.(*models.Contact)

	query := u.ContactTable().Create(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Update : Update contact to db
func (u *ContactRepository) Update(model interface{}, data map[string]interface{}) (*gorm.DB, error) {
	record := model.(*models.Contact)

	query := u.ContactTable().Model(record).Updates(data)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Destroy : destroy contact
func (u *ContactRepository) Destroy(model interface{}) (*gorm.DB, error) {
	record := model.(*models.Contact)

	query := u.ContactTable().Delete(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}
