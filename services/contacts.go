package services

import (
	"github.com/michaelt0520/nfc-card/repositories"
)

type ContactService struct {
	contactRepo *repositories.ContactRepository
}

func NewContactService(contactRepo *repositories.ContactRepository) *ContactService {
	return &ContactService{
		contactRepo: contactRepo,
	}
}

func (s *ContactService) Repo() *repositories.ContactRepository {
	return s.contactRepo
}
