package services

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/queries"
	"github.com/michaelt0520/nfc-card/repositories"
)

type CompanyService struct {
	compRepo *repositories.CompanyRepository
}

func NewCompanyService(compRepo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{
		compRepo: compRepo,
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
	if err := s.compRepo.Find(result, queries.BuildWhere(queries.BuildUserQueries(filter))); err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) FindMany(result *[]models.Company, filter map[string]interface{}, c *gin.Context) error {
	if err := s.compRepo.Where(result, queries.BuildWhere(queries.BuildUserQueries(filter)), queries.Paginate(c)); err != nil {
		return err
	}

	return nil
}
