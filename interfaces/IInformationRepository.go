package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type IInformationRepository interface {
	All(scopes ...func(db *gorm.DB) *gorm.DB) (*[]models.Information, error)
	Find(scopes ...func(db *gorm.DB) *gorm.DB) (*models.Information, error)
	Create(model *models.Information, scopes ...func(db *gorm.DB) *gorm.DB) error
	Update(model *models.Information, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) error
	Destroy(model *models.Information, scopes ...func(db *gorm.DB) *gorm.DB) error
}
