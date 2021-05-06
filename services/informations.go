package services

import (
	"github.com/michaelt0520/nfc-card/repositories"
)

type InformationService struct {
	infoRepo *repositories.InformationRepository
}

func NewInformationService(infoRepo *repositories.InformationRepository) *InformationService {
	return &InformationService{
		infoRepo: infoRepo,
	}
}

func (s *InformationService) Repo() *repositories.InformationRepository {
	return s.infoRepo
}
