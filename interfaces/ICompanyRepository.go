package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type ICompanyRepository interface {
	All(scopes ...func(db *gorm.DB) *gorm.DB) (*[]models.Company, error)
	Find(scopes ...func(db *gorm.DB) *gorm.DB) (*models.Company, error)
	Create(model *models.Company, scopes ...func(db *gorm.DB) *gorm.DB) error
	Update(model *models.Company, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) error
	Destroy(model *models.Company, scopes ...func(db *gorm.DB) *gorm.DB) error
}
