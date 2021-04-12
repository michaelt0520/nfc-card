package repositories

import "github.com/michaelt0520/nfc-card/models"

// InformationRepository : struct
type InformationRepository struct{}

// NewInformationRepository ...
func NewInformationRepository() *InformationRepository {
	return &InformationRepository{}
}

// Create : create info
func (u *InformationRepository) Create(info *models.Information) error {
	if err := GetDB().Where("name = ?", info.Name).FirstOrCreate(&info).Error; err != nil {
		return err
	}

	return nil
}

// Destroy : destroy category
func (repo *InformationRepository) Destroy(id uint) error {
	if err := GetDB().Delete(&models.Information{}, id).Error; err != nil {
		return err
	}

	return nil
}
