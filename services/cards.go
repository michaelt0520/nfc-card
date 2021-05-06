package services

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/queries"
	"github.com/michaelt0520/nfc-card/repositories"
)

type CardService struct {
	cardRepo *repositories.CardRepository
}

func NewCardService(cardRepo *repositories.CardRepository) *CardService {
	return &CardService{
		cardRepo: cardRepo,
	}
}

func (s *CardService) Repo() *repositories.CardRepository {
	return s.cardRepo
}

func (s *CardService) FindOne(result *models.Card, filter map[string]interface{}) error {
	if err := s.cardRepo.Find(result, queries.BuildWhere(queries.BuildCardQueries(filter))); err != nil {
		return err
	}

	return nil
}

func (s *CardService) FindMany(result *[]models.Card, filter map[string]interface{}, c *gin.Context) error {
	if err := s.cardRepo.Where(result, queries.BuildWhere(queries.BuildCardQueries(filter)), queries.Paginate(c)); err != nil {
		return err
	}

	return nil
}

func (s *CardService) FindManyWithScopes(result *[]models.Card, filter map[string]interface{}, preloadData map[string]interface{}, c *gin.Context) error {
	if err := s.cardRepo.Where(result, queries.BuildWhere(queries.BuildCardQueries(filter)), queries.BuildPreload(preloadData), queries.Paginate(c)); err != nil {
		return err
	}

	return nil
}

func (s *CardService) FindOneWithScopes(result *models.Card, filter map[string]interface{}, preloadData map[string]interface{}) error {
	if err := s.cardRepo.Find(result, queries.BuildWhere(queries.BuildCardQueries(filter)), queries.BuildPreload(preloadData)); err != nil {
		return err
	}

	return nil
}

func (s *CardService) RemoveCard(filter map[string]interface{}) error {
	var preloadData = map[string]interface{}{
		"User":    nil,
		"Company": nil,
	}

	var card models.Card
	if err := s.cardRepo.Find(&card, queries.BuildWhere(queries.BuildCardQueries(filter)), queries.BuildPreload(preloadData)); err != nil {
		return err
	}

	var cardUser models.User
	if err := s.cardRepo.CardTable().Model(&card).Association("User").Find(&cardUser); err != nil {
		return err
	}

	var cardCompany models.Company
	if err := s.cardRepo.CardTable().Model(&card).Association("Company").Find(&cardCompany); err != nil {
		return err
	}

	if err := s.cardRepo.CardTable().Model(&card).Association("User").Delete(&cardUser); err != nil {
		return err
	}

	if err := s.cardRepo.CardTable().Model(&card).Association("Company").Delete(&cardCompany); err != nil {
		return err
	}

	if err := s.cardRepo.Destroy(&card); err != nil {
		return err
	}

	return nil
}
