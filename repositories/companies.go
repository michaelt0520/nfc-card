package repositories

import (
	"github.com/michaelt0520/nfc-card/models"
)

// CompanyRepository : struct
type CompanyRepository struct{}

// NewCompanyRepository ...
func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{}
}

// Find : get company by id
func (u *CompanyRepository) Find(id uint) (*models.Company, error) {
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
func (repo *CompanyRepository) Update(record *models.Company, data map[string]interface{}) (*models.Company, error) {
	if err := GetDB().Model(&record).Updates(data).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// Destroy : destroy company
func (repo *CompanyRepository) Destroy(id uint) error {
	if err := GetDB().Delete(&models.Company{}, id).Error; err != nil {
		return err
	}

	return nil
}
