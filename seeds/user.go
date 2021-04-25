package seeds

import (
	"fmt"
	"os"

	"github.com/michaelt0520/nfc-card/models"
	"github.com/michaelt0520/nfc-card/repositories"
	"go.uber.org/zap"
)

const DefaultAvatar = "/public/avatars/default_avatar.png"

// LoadUserSeed : constructor user seed data
func (seed *Seed) LoadUserSeed() {
	repoUser := repositories.NewUserRepository()

	users := []models.User{
		{Name: "1TAP First", Username: "1tap01", Email: "mt520@1tap.vn", Password: "haituan", Address: "76 Wall St", PhoneNumber: "0909 990 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(1), Role: models.UserAdmin},
		{Name: "1TAP Second", Username: "1tap02", Email: "mt0520@1tap.vn", Password: "haituan", Address: "78 Wall St", PhoneNumber: "0909 991 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyManager},
		{Name: "1TAP Third", Username: "1tap03", Email: "mt0521@1tap.vn", Password: "haituan", Address: "80 Wall St", PhoneNumber: "0909 992 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "1TAP Fourth", Username: "1tap04", Email: "mt0522@1tap.vn", Password: "haituan", Address: "82 Wall St", PhoneNumber: "0909 993 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(2), Role: models.UserCompanyManager},
		{Name: "1TAP Fifth", Username: "1tap05", Email: "mt0523@1tap.vn", Password: "haituan", Address: "84 Wall St", PhoneNumber: "0909 994 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(2), Role: models.UserCompanyMember},
		{Name: "1TAP Sixth", Username: "1tap06", Email: "mt0524@1tap.vn", Password: "haituan", Address: "86 Wall St", PhoneNumber: "0909 995 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(2), Role: models.UserCompanyMember},
		{Name: "1TAP Seventh", Username: "1tap07", Email: "mt0525@1tap.vn", Password: "haituan", Address: "88 Wall St", PhoneNumber: "0909 996 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(3), Role: models.UserCompanyManager},
		{Name: "1TAP Eighth", Username: "1tap08", Email: "mt0526@1tap.vn", Password: "haituan", Address: "90 Wall St", PhoneNumber: "0909 997 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(3), Role: models.UserCompanyMember},
		{Name: "1TAP Nineth", Username: "1tap09", Email: "mt0527@1tap.vn", Password: "haituan", Address: "92 Wall St", PhoneNumber: "0909 998 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, CompanyID: models.ConvertPointerUint(3), Role: models.UserCompanyMember},
		{Name: "1TAP Tenth", Username: "1tap10", Email: "mt0528@1tap.vn", Password: "haituan", Address: "94 Wall St", PhoneNumber: "0909 999 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
		{Name: "1TAP Eleventth", Username: "1tap11", Email: "mt0529@1tap.vn", Password: "haituan", Address: "96 Wall St", PhoneNumber: "0910 000 009", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
	}

	for _, user := range users {
		err := repoUser.Create(&user)
		if err != nil {
			seed.log.Error("failed to load seed: User", zap.Error(err))
		}
	}
}
