package seeds

import (
	"fmt"
	"os"

	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"go.uber.org/zap"
)

const DefaultAvatar = "/public/images/default_avatar.png"

// LoadUserSeed : constructor user seed data
func (seed *Seed) LoadUserSeed() {
	repoUser := repositories.NewUserRepository()

	users := []models.User{
		{Name: "1TAP First", Username: "1tap01", Email: "mt520@1tap.vn", Password: "haituan", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: 1},
		{Name: "1TAP Second", Username: "1tap02", Email: "mt0520@1tap.vn", Password: "haituan", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: 1},
		{Name: "1TAP Third", Username: "1tap03", Email: "mt0521@1tap.vn", Password: "haituan", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: 2},
		{Name: "1TAP Fourth", Username: "1tap04", Email: "mt0522@1tap.vn", Password: "haituan", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: 2},
		{Name: "1TAP Fifth", Username: "1tap05", Email: "mt0523@1tap.vn", Password: "haituan", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: 3},
		{Name: "1TAP Sixth", Username: "1tap06", Email: "mt0524@1tap.vn", Password: "haituan", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: 3},
	}

	for _, user := range users {
		err := repoUser.Create(&user)
		if err != nil {
			seed.log.Error("failed to load seed: User", zap.Error(err))
		}
	}
}
