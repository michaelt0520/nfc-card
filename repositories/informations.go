package repositories

import (
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

var _ Repository = &InformationRepository{}

// Find : get card by code
func (u *InformationRepository) All(result interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.InformationTable().Scopes(scopes...).Preload("User").Find(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Find : get info by id
func (u *InformationRepository) Find(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.InformationTable().Scopes(scopes...).Preload("User").Where(data).First(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Where :
func (u *InformationRepository) Where(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.InformationTable().Scopes(scopes...).Preload("User").Where(data).Find(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Create : Save user to db
func (u *InformationRepository) Create(model interface{}) (*gorm.DB, error) {
	record := model.(*models.Information)

	query := u.InformationTable().Create(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Update : Update info to db
func (u *InformationRepository) Update(model interface{}, data map[string]interface{}) (*gorm.DB, error) {
	record := model.(*models.Information)

	query := u.InformationTable().Model(record).Updates(data)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Destroy : destroy info
func (u *InformationRepository) Destroy(model interface{}) (*gorm.DB, error) {
	record := model.(*models.Information)

	query := u.InformationTable().Delete(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}
