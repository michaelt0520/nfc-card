package repositories

import (
	"github.com/michaelt0520/nfc-card/interfaces"
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

var _ interfaces.ICompanyRepository = &CompanyRepository{}

// Find : get company by id
func (u *CompanyRepository) Find(result *models.Company, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.CompanyTable().Scopes(scopes...).First(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Where :
func (u *CompanyRepository) Where(result *[]models.Company, scopes ...func(db *gorm.DB) *gorm.DB) error {
	query := u.CompanyTable().Scopes(scopes...).Find(result)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Create : create company
func (u *CompanyRepository) Create(model *models.Company) error {
	query := u.CompanyTable().Where(models.Company{Name: model.Name}).FirstOrCreate(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Update : Update company to db
func (u *CompanyRepository) Update(model *models.Company, data map[string]interface{}) error {
	query := u.CompanyTable().Model(model).Updates(data)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// Destroy : destroy company
func (u *CompanyRepository) Destroy(model *models.Company) error {
	query := u.CompanyTable().Delete(model)
	if err := query.Error; err != nil {
		return err
	}

	return nil
}
