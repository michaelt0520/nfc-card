package repositories

import (
	"fmt"
	"strings"

	"github.com/michaelt0520/nfc-card/models"
	"github.com/mitchellh/mapstructure"
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

var _ Repository = &CardRepository{}

// Find : get card by code
func (u *CardRepository) All(result interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.CardTable().Scopes(scopes...).Preload("User").Find(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Find : get card by code
func (u *CardRepository) Find(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.CardTable().Scopes(scopes...).Preload("User.Informations").Preload("User.Company").Where(data).First(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Where :
func (u *CardRepository) Where(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	var fields models.Card
	mapstructure.Decode(data, &fields)

	query := u.CardTable().Scopes(scopes...).Preload("User.Informations").Preload("User.Company").Where(fields, "company_id")
	if err := query.Error; err != nil {
		return nil, err
	}

	if data["code"] != nil {
		code := fmt.Sprintf("%v", data["code"])
		query.Where("lower(code) LIKE ?", fmt.Sprintf("%%%v%%", strings.ToLower(code)))
	}

	return query.Find(result), nil
}

// Create : Save card to db
func (u *CardRepository) Create(model interface{}) (*gorm.DB, error) {
	record := model.(*models.Card)
	query := u.CardTable().Where("code = ?", record.Code).FirstOrCreate(&record)
	if err := query.Error; err != nil {
		return query, nil
	}

	return query, nil
}

// Update : Update card to db
func (u *CardRepository) Update(model interface{}, data map[string]interface{}) (*gorm.DB, error) {
	record := model.(*models.Card)

	query := u.CardTable().Model(record).Updates(data)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Destroy : destroy card
func (u *CardRepository) Destroy(model interface{}) (*gorm.DB, error) {
	record := model.(*models.Card)

	query := u.CardTable().Delete(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}
