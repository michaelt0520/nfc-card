package services

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/queries"
	"github.com/michaelt0520/nfc-card/repositories"
)

type CompanyService struct {
	compRepo *repositories.CompanyRepository
	userRepo *repositories.UserRepository
}

func NewCompanyService(compRepo *repositories.CompanyRepository, userRepo *repositories.UserRepository) *CompanyService {
	return &CompanyService{
		compRepo: compRepo,
		userRepo: userRepo,
	}
}

func (s *CompanyService) Repo() *repositories.CompanyRepository {
	return s.compRepo
}

func (s *CompanyService) GetCurrentCompany(c *gin.Context) (*models.Company, error) {
	companyID := c.MustGet("company_id")

	var companyFilter = map[string]interface{}{
		"id = ?": companyID,
	}

	var company models.Company
	if err := s.compRepo.Find(&company, queries.BuildWhere(companyFilter)); err != nil {
		return nil, err
	}

	return &company, nil
}

func (s *CompanyService) FindOne(result *models.Company, filter map[string]interface{}) error {
	if err := s.compRepo.Find(result, queries.BuildWhere(queries.BuildFinds(filter))); err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) FindMany(result *[]models.Company, filter map[string]interface{}, c *gin.Context) error {
	if err := s.compRepo.Where(result, queries.BuildWhere(queries.BuildQueries(filter)), queries.Paginate(c)); err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) AddUserToCompany(company *models.Company, user *models.User) error {
	if err := s.compRepo.CompanyTable().Model(company).Association("Users").Append(user); err != nil {
		return err
	}

	var updateData = map[string]interface{}{
		"role": models.UserCompanyMember,
		"type": models.Business,
	}

	if err := s.userRepo.Update(user, updateData); err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) RemoveUserFromCompany(company *models.Company, user *models.User) error {
	var userCompanyCards []models.Card
	if err := s.compRepo.CompanyTable().Model(company).Association("Cards").Find(&userCompanyCards); err != nil {
		return err
	}

	if err := s.userRepo.UserTable().Model(user).Association("Cards").Delete(userCompanyCards); err != nil {
		return err
	}

	if err := s.compRepo.CompanyTable().Model(company).Association("Users").Delete(user); err != nil {
		return err
	}

	var updateData = map[string]interface{}{
		"role": models.UserStandard,
		"type": models.Personal,
	}

	if err := s.userRepo.Update(user, updateData); err != nil {
		return err
	}

	return nil
}
