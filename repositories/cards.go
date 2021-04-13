package repositories

import (
	"github.com/michaelt0520/nfc-card/models"
)

// CardRepository : struct
type CardRepository struct{}

// NewCardRepository ...
func NewCardRepository() *CardRepository {
	return &CardRepository{}
}

// Find : get card by code
func (u *CardRepository) All() *[]models.Card {
  var cards []models.Card
	GetDB().Find(&cards)

  return &cards
}

// Find : get card by code
func (u *CardRepository) Find(cardCode string) (*models.Card, error) {
	var card models.Card

	if err := GetDB().Where("code = ?", cardCode).First(&card).Error; err != nil {
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
func (repo *CardRepository) Update(record *models.Card, data map[string]interface{}) (*models.Card, error) {
	if err := GetDB().Model(&record).Updates(data).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// Destroy : destroy card
func (repo *CardRepository) Destroy(cardCode string) error {
	if err := GetDB().Where("code = ?", cardCode).Delete(&models.Card{}).Error; err != nil {
		return err
	}

	return nil
}
