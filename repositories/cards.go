package repositories

import (
	"github.com/michaelt0520/nfc-card/interfaces"
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

// CardRepository : struct
type CardRepository struct {
}

// NewCardRepository ...
func NewCardRepository() *CardRepository {
	return &CardRepository{}
}

func (u *CardRepository) CardTable() *gorm.DB {
	return GetDB().Model(&models.Card{})
}

var _ interfaces.ICardRepository = &CardRepository{}

// Find : get card by code
func (u *CardRepository) Find(result *models.Card, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.CardTable().Scopes(scopes...).First(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Where :
func (u *CardRepository) Where(result *[]models.Card, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.CardTable().Scopes(scopes...).Find(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Create : Save card to db
func (u *CardRepository) Create(model *models.Card) error {
	query := u.CardTable().Where("code = ?", model.Code).FirstOrCreate(&model)
	if err := query.Error; err != nil {
		return nil
	}

	return nil
}

// Update : Update card to db
func (u *CardRepository) Update(model *models.Card, data map[string]interface{}) error {
	query := u.CardTable().Model(model).Updates(data)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Destroy : destroy card
func (u *CardRepository) Destroy(model *models.Card) error {
	query := u.CardTable().Delete(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}
