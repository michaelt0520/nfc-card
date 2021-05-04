package repositories

import (
	"github.com/michaelt0520/nfc-card/models"
	"gorm.io/gorm"
)

// CompanyRepository : struct
type CompanyRepository struct{}

// NewCompanyRepository ...
func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{}
}

func (u *CompanyRepository) CompanyTable() *gorm.DB {
	return GetDB().Model(&models.Company{})
}

var _ Repository = &CompanyRepository{}

// Find : get card by code
func (u *CompanyRepository) All(result interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.CompanyTable().Scopes(scopes...).Find(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Find : get company by id
func (u *CompanyRepository) Find(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.CompanyTable().Scopes(scopes...).Preload("Users").Where(data).First(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Where :
func (u *CompanyRepository) Where(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error) {
	query := u.CompanyTable().Scopes(scopes...).Preload("Users").Where(data).Find(result)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Create : create company
func (u *CompanyRepository) Create(model interface{}) (*gorm.DB, error) {
	record := model.(*models.Company)
	query := u.CompanyTable().Where(models.Company{Name: record.Name}).FirstOrCreate(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Update : Update company to db
func (u *CompanyRepository) Update(model interface{}, data map[string]interface{}) (*gorm.DB, error) {
	record := model.(*models.Company)

	query := u.CompanyTable().Model(record).Updates(data)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}

// Destroy : destroy company
func (u *CompanyRepository) Destroy(model interface{}) (*gorm.DB, error) {
	record := model.(*models.Company)

	query := u.CompanyTable().Delete(record)
	if err := query.Error; err != nil {
		return nil, err
	}

	return query, nil
}
