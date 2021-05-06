package repositories

import (
	"github.com/michaelt0520/nfc-card/interfaces"
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

var _ interfaces.IContactRepository = &ContactRepository{}

// Find : get contact by id
func (u *ContactRepository) Find(result *models.Contact, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.ContactTable().Scopes(scopes...).First(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Create : create contact
func (u *ContactRepository) Where(result *[]models.Contact, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.ContactTable().Scopes(scopes...).Find(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Create : Save user to db
func (u *ContactRepository) Create(model *models.Contact) error {
	query := u.ContactTable().Create(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Update : Update contact to db
func (u *ContactRepository) Update(model *models.Contact, data map[string]interface{}) error {
	query := u.ContactTable().Model(model).Updates(data)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Destroy : destroy contact
func (u *ContactRepository) Destroy(model *models.Contact) error {
	query := u.ContactTable().Delete(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}
