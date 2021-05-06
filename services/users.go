package services

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/queries"
	"github.com/michaelt0520/nfc-card/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Repo() *repositories.UserRepository {
	return s.userRepo
}

func (s *UserService) GetCurrentUser(c *gin.Context) (*models.User, error) {
	userID := c.MustGet("user_id")

	var filterUser = map[string]interface{}{
		"id = ?": userID,
	}

	var user models.User
	if err := s.userRepo.Find(&user, queries.BuildWhere(filterUser)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) FindOne(result *models.User, filter map[string]interface{}) error {
	if err := s.userRepo.Find(result, queries.BuildWhere(queries.BuildUserQueries(filter))); err != nil {
		return err
	}

	return nil
}

func (s *UserService) FindOneWithScopes(result *models.User, filter map[string]interface{}, preloadData map[string]interface{}) error {
	if err := s.userRepo.Find(result, queries.BuildWhere(queries.BuildUserQueries(filter)), queries.BuildPreload(preloadData)); err != nil {
		return err
	}

	return nil
}

func (s *UserService) FindMany(result *[]models.User, filter map[string]interface{}) error {
	if err := s.userRepo.Where(result, queries.BuildWhere(queries.BuildUserQueries(filter))); err != nil {
		return err
	}

	return nil
}

func (s *UserService) AddInformationToUser(user *models.User, info *models.Information) error {
	if err := s.userRepo.UserTable().Model(&user).Association("Informations").Append(&info); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetInformationFromUser(user *models.User, id string, result *models.Information) error {
	if err := s.userRepo.UserTable().Model(&user).Association("Informations").Find(&result, id); err != nil {
		return err
	}

	return nil
}
