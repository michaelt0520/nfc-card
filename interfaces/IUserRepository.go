package interfaces

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Find(result *models.User, scopes ...func(db *gorm.DB) *gorm.DB) error
	Where(result *[]models.User, scopes ...func(db *gorm.DB) *gorm.DB) error
	Create(model *models.User) error
	Update(model *models.User, data map[string]interface{}) error
	Destroy(model *models.User) error
}
