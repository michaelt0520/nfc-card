package seeds

import (
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"go.uber.org/zap"
)

// LoadCardSeed : constructor card seed data
func (seed *Seed) LoadCardSeed() {
	cardRepo := repositories.NewCardRepository()

	cards := []models.Card{
		{Code: "ONETAP001", Activated: true, UserID: models.ConvertPointerUint(1), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP002", Activated: true, UserID: models.ConvertPointerUint(2), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP003", Activated: true, UserID: models.ConvertPointerUint(3), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP004", Activated: true, UserID: models.ConvertPointerUint(4), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP005", Activated: true, UserID: models.ConvertPointerUint(5), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP006", Activated: true, UserID: models.ConvertPointerUint(6), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP007", Activated: true, UserID: models.ConvertPointerUint(7), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP008", Activated: true, UserID: models.ConvertPointerUint(8), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP009", Activated: true, UserID: models.ConvertPointerUint(9), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP010", Activated: true, UserID: models.ConvertPointerUint(10), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP011", Activated: true, UserID: models.ConvertPointerUint(11), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP012", Activated: true, UserID: models.ConvertPointerUint(12), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP013", Activated: true, UserID: models.ConvertPointerUint(13), CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP014", Activated: true, CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP015", Activated: true, CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP016", Activated: true, CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP017", Activated: true, CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP018", Activated: true, CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP019", Activated: true, CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP020", Activated: true, CompanyID: models.ConvertPointerUint(1)},
		{Code: "ONETAP021", Activated: true, CompanyID: models.ConvertPointerUint(2)},
		{Code: "ONETAP022", Activated: true, CompanyID: models.ConvertPointerUint(2)},
		{Code: "ONETAP023", Activated: true},
		{Code: "ONETAP024", Activated: true},
		{Code: "ONETAP025", Activated: true},
		{Code: "ONETAP026", Activated: true},
		{Code: "ONETAP027", Activated: true},
		{Code: "ONETAP028", Activated: true},
		{Code: "ONETAP029", Activated: true},
		{Code: "ONETAP030", Activated: true},
	}

	for _, card := range cards {
		if err := cardRepo.Create(&card); err != nil {
			seed.log.Error("failed to load seed: Card", zap.Error(err))
		}
	}
}
