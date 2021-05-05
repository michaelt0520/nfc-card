package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	All(scopes ...func(db *gorm.DB) *gorm.DB) (*[]models.User, error)
	Find(scopes ...func(db *gorm.DB) *gorm.DB) (*models.User, error)
	Create(model *models.User, scopes ...func(db *gorm.DB) *gorm.DB) error
	Update(model *models.User, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) error
	Destroy(model *models.User, scopes ...func(db *gorm.DB) *gorm.DB) error
}
