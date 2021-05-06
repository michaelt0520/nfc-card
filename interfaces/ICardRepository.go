package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type ICardRepository interface {
	Find(result *models.Card, scopes ...func(db *gorm.DB) *gorm.DB) error
	Where(result *[]models.Card, scopes ...func(db *gorm.DB) *gorm.DB) error
	Create(model *models.Card) error
	Update(model *models.Card, data map[string]interface{}) error
	Destroy(model *models.Card) error
}
