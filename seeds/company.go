package seeds

import (
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"go.uber.org/zap"
)

// LoadCompanySeed : constructor company seed data
func (seed *Seed) LoadCompanySeed() {
	compRepo := repositories.NewCompanyRepository()

	companies := []models.Company{
		{Name: "1PAC Vietnam Co.,Ltd", Address: "5th Floor, Devspace Building, 86-88 Dinh Tien Hoang, Ward 1, Binh Thanh District, Ho Chi Minh City, Vietnam", Logo: "https://1pac.vn/img/common/loading.gif", Website: "https://1pac.vn/"},
		{Name: "FPT Software", Address: "Ho Chi Minh, Ha Noi, Da Nang", Logo: "https://employer.jobsgo.vn/uploads/media/img/201911/pictures_library_24959_20191116115153_9441.jpg", Website: "https://www.fpt-software.com/"},
		{Name: "Toshiba Software Development", Address: "Ba Dinh, Ha Noi", Logo: "https://upload.wikimedia.org/wikipedia/commons/thumb/0/01/Toshiba_logo.svg/1280px-Toshiba_logo.svg.png", Website: "https://www.toshiba.com.vn/"},
	}

	for _, company := range companies {
		err := compRepo.Create(&company)
		if err != nil {
			seed.log.Error("failed to load seed: Company", zap.Error(err))
		}
	}
}
