package repositories

import (
	"github.com/michaelt0520/nfc-card/models"
)

// ContactRepository : struct
type ContactRepository struct{}

// NewContactRepository ...
func NewContactRepository() *ContactRepository {
	return &ContactRepository{}
}

// Find : get contact by id
func (u *ContactRepository) Find(data map[string]interface{}) (*models.Contact, error) {
	var contact models.Contact

	if err := GetDB().Where(data).First(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

// Create : create contact
func (u *ContactRepository) Create(contact *models.Contact) error {
	if err := GetDB().Create(&contact).Error; err != nil {
		return err
	}

	return nil
}

// Update : Update contact to db
func (repo *ContactRepository) Update(record *models.Contact, data map[string]interface{}) error {
	if err := GetDB().Model(&record).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// Destroy : destroy contact
func (repo *ContactRepository) Destroy(record *models.Contact) error {
	if err := GetDB().Delete(&record).Error; err != nil {
		return err
	}

	return nil
}
