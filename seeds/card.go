package seeds

import (
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"go.uber.org/zap"
)

// LoadCardSeed : constructor card seed data
func (seed *Seed) LoadCardSeed() {
	cardRepo := repositories.CardRepository{}

	cards := []models.Card{
		{Code: "1TAP001", Activated: true, UserID: 1, CompanyID: 1},
		{Code: "1TAP002", Activated: true, UserID: 1, CompanyID: 1},
		{Code: "1TAP003", Activated: true, UserID: 1, CompanyID: 1},
		{Code: "1TAP004", Activated: true, UserID: 1, CompanyID: 1},
		{Code: "1TAP005", Activated: true, UserID: 2, CompanyID: 1},
		{Code: "1TAP006", Activated: true, UserID: 2, CompanyID: 1},
		{Code: "1TAP007", Activated: true, UserID: 2, CompanyID: 2},
		{Code: "1TAP008", Activated: true, UserID: 2, CompanyID: 2},
		{Code: "1TAP009", Activated: true, UserID: 3, CompanyID: 2},
		{Code: "1TAP010", Activated: true, UserID: 3, CompanyID: 2},
		{Code: "1TAP011", Activated: true, UserID: 3, CompanyID: 2},
		{Code: "1TAP012", Activated: true, UserID: 3, CompanyID: 2},
		{Code: "1TAP013", Activated: true, UserID: 4, CompanyID: 2},
		{Code: "1TAP014", Activated: true, UserID: 4, CompanyID: 3},
		{Code: "1TAP015", Activated: true, UserID: 4, CompanyID: 3},
		{Code: "1TAP016", Activated: true, UserID: 5, CompanyID: 3},
		{Code: "1TAP017", Activated: true, UserID: 5, CompanyID: 3},
		{Code: "1TAP018", Activated: true, UserID: 5, CompanyID: 3},
		{Code: "1TAP019", Activated: true, UserID: 6, CompanyID: 3},
		{Code: "1TAP020", Activated: true, UserID: 6, CompanyID: 3},
	}

	for _, card := range cards {
		err := cardRepo.Create(&card)
		if err != nil {
			seed.log.Error("failed to load seed: Card", zap.Error(err))
		}
	}
}
