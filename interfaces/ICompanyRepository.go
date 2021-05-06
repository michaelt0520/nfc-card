package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type ICompanyRepository interface {
	Find(result *models.Company, scopes ...func(db *gorm.DB) *gorm.DB) error
	Where(result *[]models.Company, scopes ...func(db *gorm.DB) *gorm.DB) error
	Create(model *models.Company) error
	Update(model *models.Company, data map[string]interface{}) error
	Destroy(model *models.Company) error
}
