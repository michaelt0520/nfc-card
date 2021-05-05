package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type ICardRepository interface {
	All(scopes ...func(db *gorm.DB) *gorm.DB) (*[]models.Card, error)
	Find(scopes ...func(db *gorm.DB) *gorm.DB) (*models.Card, error)
	Create(model *models.Card, scopes ...func(db *gorm.DB) *gorm.DB) error
	Update(model *models.Card, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) error
	Destroy(model *models.Card, scopes ...func(db *gorm.DB) *gorm.DB) error
}
