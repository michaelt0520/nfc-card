package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type IInformationRepository interface {
	Find(result *models.Information, scopes ...func(db *gorm.DB) *gorm.DB) error
	Where(result *[]models.Information, scopes ...func(db *gorm.DB) *gorm.DB) error
	Create(model *models.Information) error
	Update(model *models.Information, data map[string]interface{}) error
	Destroy(model *models.Information) error
}
