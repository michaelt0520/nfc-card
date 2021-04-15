package repositories

import (
	"github.com/michaelt0520/nfc-card/models"
)

// InformationRepository : struct
type InformationRepository struct{}

// NewInformationRepository ...
func NewInformationRepository() *InformationRepository {
	return &InformationRepository{}
}

// Find : get info by id
func (u *InformationRepository) Find(id uint) (*models.Information, error) {
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
func (repo *InformationRepository) Update(record *models.Information, data map[string]interface{}) error {
	if err := GetDB().Model(&record).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// Destroy : destroy info
func (repo *InformationRepository) Destroy(id uint) error {
	if err := GetDB().Delete(&models.Information{}, id).Error; err != nil {
		return err
	}

	return nil
}
