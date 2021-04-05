package repositories

import "github.com/michaelt0520/nfc-card/models"

// SocialInformationRepository : struct
type SocialInformationRepository struct{}

// NewSocialInformationRepository ...
func NewSocialInformationRepository() *SocialInformationRepository {
	return &SocialInformationRepository{}
}

// Create : create social info
func (u *SocialInformationRepository) Create(username string) (*models.SocialInformation, error) {
	var social_info models.SocialInformation

	if err := GetDB().Where("username = ?", username).Find(&social_info).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
