package repositories

import (
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
)

// InformationRepository : struct
type InformationRepository struct{}

// NewInformationRepository ...
func NewInformationRepository() *InformationRepository {
	return &InformationRepository{}
}

// Find : get user by username
func (u *InformationRepository) Find(id interface{}) (*models.Information, error) {
	var info models.Information

	if err := GetDB().First(&info, id).Error; err != nil {
		return nil, err
	}

	return &info, nil
}

// Create : create info
func (u *InformationRepository) Create(info *models.Information) error {
	if err := GetDB().Where("name = ?", info.Name).FirstOrCreate(&info).Error; err != nil {
		return err
	}

	return nil
}

// Update : Update info to db
func (repo *InformationRepository) Update(data map[string]interface{}) (*models.Information, error) {
	info, err := repo.Find(data["id"])
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.RecordNotFound
	}

	if err := GetDB().Model(&info).Where("id = ?", data["id"]).Updates(data).Error; err != nil {
		return nil, err
	}

	return info, nil
}

// Destroy : destroy category
func (repo *InformationRepository) Destroy(id uint) error {
	if err := GetDB().Delete(&models.Information{}, id).Error; err != nil {
		return err
	}

	return nil
}
