package repositories

import (
	"github.com/michaelt0520/nfc-card/interfaces"
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

// InformationRepository : struct
type InformationRepository struct{}

// NewInformationRepository ...
func NewInformationRepository() *InformationRepository {
	return &InformationRepository{}
}

func (u *InformationRepository) InformationTable() *gorm.DB {
	return GetDB().Model(&models.Information{})
}

var _ interfaces.IInformationRepository = &InformationRepository{}

// Find : get info by id
func (u *InformationRepository) Find(result *models.Information, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.InformationTable().Scopes(scopes...).First(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Where :
func (u *InformationRepository) Where(result *[]models.Information, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.InformationTable().Scopes(scopes...).Find(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Create : Save user to db
func (u *InformationRepository) Create(model *models.Information) error {
	query := u.InformationTable().Create(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Update : Update info to db
func (u *InformationRepository) Update(model *models.Information, data map[string]interface{}) error {
	query := u.InformationTable().Model(model).Updates(data)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Destroy : destroy info
func (u *InformationRepository) Destroy(model *models.Information) error {
	query := u.InformationTable().Delete(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}
