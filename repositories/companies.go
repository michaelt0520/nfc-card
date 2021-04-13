package repositories

import (
	"github.com/michaelt0520/nfc-card/errors"
	"github.com/michaelt0520/nfc-card/models"
)

// CompanyRepository : struct
type CompanyRepository struct{}

// NewCompanyRepository ...
func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{}
}

// Find : get user by username
func (u *CompanyRepository) Find(id interface{}) (*models.Company, error) {
	var company models.Company

	if err := GetDB().First(&company, id).Error; err != nil {
		return nil, err
	}

	return &company, nil
}

// Create : create company
func (u *CompanyRepository) Create(company *models.Company) error {
	if err := GetDB().Where("name = ?", company.Name).FirstOrCreate(&company).Error; err != nil {
		return err
	}

	return nil
}

// Update : Update company to db
func (repo *CompanyRepository) Update(data map[string]interface{}) (*models.Company, error) {
	company, err := repo.Find(data["id"])
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, errors.RecordNotFound
	}

	if err := GetDB().Model(&company).Where("id = ?", data["id"]).Updates(data).Error; err != nil {
		return nil, err
	}

	return company, nil
}

// Destroy : destroy category
func (repo *CompanyRepository) Destroy(id uint) error {
	if err := GetDB().Delete(&models.Company{}, id).Error; err != nil {
		return err
	}

	return nil
}
