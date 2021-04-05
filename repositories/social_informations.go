package repositories

import "github.com/michaelt0520/nfc-card/models"

// SocialInformationRepository : struct
type SocialInformationRepository struct{}

// NewSocialInformationRepository ...
func NewSocialInformationRepository() *SocialInformationRepository {
	return &SocialInformationRepository{}
}

// Create : create social info
func (u *SocialInformationRepository) Create(socialInfo *models.SocialInformation) error {
	if err := GetDB().Create(&socialInfo).Error; err != nil {
		return err
	}

	return nil
}

// Destroy : destroy category
func (repo *SocialInformationRepository) Destroy(id uint) error {
	if err := GetDB().Delete(&models.SocialInformation{}, id).Error; err != nil {
		return err
	}

	return nil
}