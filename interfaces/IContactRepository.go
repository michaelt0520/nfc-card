package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type IContactRepository interface {
	Find(result *models.Contact, scopes ...func(db *gorm.DB) *gorm.DB) error
	Where(result *[]models.Contact, scopes ...func(db *gorm.DB) *gorm.DB) error
	Create(model *models.Contact) error
	Update(model *models.Contact, data map[string]interface{}) error
	Destroy(model *models.Contact) error
}
