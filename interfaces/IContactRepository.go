package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type IContactRepository interface {
	All(scopes ...func(db *gorm.DB) *gorm.DB) (*[]models.Contact, error)
	Find(scopes ...func(db *gorm.DB) *gorm.DB) (*models.Contact, error)
	Create(model *models.Contact, scopes ...func(db *gorm.DB) *gorm.DB) error
	Update(model *models.Contact, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) error
	Destroy(model *models.Contact, scopes ...func(db *gorm.DB) *gorm.DB) error
}
