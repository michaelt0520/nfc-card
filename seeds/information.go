package seeds

import (
	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"go.uber.org/zap"
)

// LoadInformationSeed : constructor card seed data
func (seed *Seed) LoadInformationSeed() {
	infoRepo := repositories.NewInformationRepository()

	infos := []models.Information{
		{Name: "Facebook", Data: "https://facebook.com", Visibled: true, UserID: 1},
		{Name: "Instagram", Data: "https://instagram.com", Visibled: true, UserID: 1},
		{Name: "Youtube", Data: "https://youtube.com", Visibled: true, UserID: 1},
		{Name: "WhatsApp", Data: "https://whatsapp.com", Visibled: true, UserID: 1},
		{Name: "TikTok", Data: "https://tiktok.com", Visibled: true, UserID: 2},
		{Name: "WeChat", Data: "https://wechat.com", Visibled: true, UserID: 2},
		{Name: "Tumblr", Data: "https://tumblr.com", Visibled: true, UserID: 2},
		{Name: "Qzone", Data: "https://qzone.qq.com", Visibled: true, UserID: 2},
		{Name: "QQ", Data: "https://imqq.com", Visibled: true, UserID: 3},
		{Name: "Twitter", Data: "https://twitter.com", Visibled: true, UserID: 3},
		{Name: "Reddit", Data: "https://reddit.com", Visibled: true, UserID: 3},
		{Name: "LinkedIn", Data: "https://linkedin.com", Visibled: true, UserID: 3},
		{Name: "GitHub", Data: "https://github.com", Visibled: true, UserID: 4},
		{Name: "Viber", Data: "https://viber.com", Visibled: true, UserID: 4},
		{Name: "Snapchat", Data: "https://snapchat.com", Visibled: true, UserID: 4},
		{Name: "Pinterest", Data: "https://pinterest.com", Visibled: true, UserID: 5},
		{Name: "Line", Data: "https://line.me", Visibled: true, UserID: 5},
		{Name: "Telegram", Data: "https://telegram.org", Visibled: true, UserID: 5},
		{Name: "Medium", Data: "https://medium.com", Visibled: true, UserID: 6},
		{Name: "Zalo", Data: "https://zalo.me", Visibled: true, UserID: 6},
	}

	for _, info := range infos {
		if err := infoRepo.Create(&info); err != nil {
			seed.log.Error("failed to load seed: Information", zap.Error(err))
		}
	}
}
