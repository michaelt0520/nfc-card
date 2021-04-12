package repositories

import "github.com/michaelt0520/nfc-card/models"

// CardRepository : struct
type CardRepository struct{}

// NewCardRepository ...
func NewCardRepository() *CardRepository {
	return &CardRepository{}
}

// Find : get card by code
func (u *CardRepository) Find(cardCode string) (*models.Card, error) {
	var card models.Card

	if err := GetDB().Where("code = ?", cardCode).Find(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}

// Create : Save card to db
func (repo *CardRepository) Create(card *models.Card) error {
	if err := GetDB().Where("code = ?", card.Code).FirstOrCreate(&card).Error; err != nil {
		return err
	}

	return nil
}

// Update : Update card to db
func (repo *CardRepository) Update(data map[string]interface{}) (*models.Card, error) {
	var card models.Card
	if err := GetDB().Model(&card).Where("code = ?", data["code"]).Updates(data).Error; err != nil {
		return nil, err
	}

	return &card, nil
}
