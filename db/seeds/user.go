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
		{Name: "Admin", Username: "onetap01", Email: "admin@1pac.vn", Password: "onetap", Address: "74 Wall St", PhoneNumber: "0909 664 641", Avatar: "https://1pac.vn/img/common/loading.gif", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserAdmin},
		{Name: "Michael", Username: "onetap02", Email: "michael.tuan@1pac.vn", Password: "onetap", Address: "76 Wall St", PhoneNumber: "0909 664 641", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-UM0LKB40N-f9839bd17c1c-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyManager},
		{Name: "Nghia Tran", Username: "onetap03", Email: "tran.nghia@1pac.vn", Password: "onetap", Address: "78 Wall St", PhoneNumber: "0909 991 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U014Y9J08EL-b845fd27c897-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Quang Vu", Username: "onetap04", Email: "vv.quang@1pac.vn", Password: "onetap", Address: "80 Wall St", PhoneNumber: "0909 992 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U01278981AT-7d283bb1c16b-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Vinh Kieu", Username: "onetap05", Email: "kq.vinh@1pac.vn", Password: "onetap", Address: "82 Wall St", PhoneNumber: "0909 993 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U01133DB52R-de7ac1e638f9-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Vinh Mai", Username: "onetap06", Email: "vinhmai570@gmail.com", Password: "onetap", Address: "84 Wall St", PhoneNumber: "0909 994 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U01JCLH2SHE-4d46bc82e146-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Dung Ga", Username: "onetap07", Email: "ga.dung@1pac.vn", Password: "onetap", Address: "86 Wall St", PhoneNumber: "0909 995 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U01E5HUL2AD-6fe371cdbf42-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Khang Huynh", Username: "onetap08", Email: "lam.khang@1pac.vn", Password: "onetap", Address: "88 Wall St", PhoneNumber: "0909 995 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U01TD7W4H41-aeae96e26a8d-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Nhat My", Username: "onetap09", Email: "nhat.my@1pac.vn", Password: "onetap", Address: "90 Wall St", PhoneNumber: "0909 996 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U5PP5K3SL-6b47319e9ec4-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Kusa Thao", Username: "onetap10", Email: "kusa.thao@1pac.vn", Password: "onetap", Address: "92 Wall St", PhoneNumber: "0909 997 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-UFRS5B39C-669d3be3706c-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Fu Aya", Username: "onetap11", Email: "fu.aya@1pac.vn", Password: "onetap", Address: "94 Wall St", PhoneNumber: "0909 998 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-USSMQ7SMT-5ce3ed34666d-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Thao Ha", Username: "onetap12", Email: "hp.thao@1pac.vn", Password: "onetap", Address: "96 Wall St", PhoneNumber: "0909 999 009", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U4W22T1JL-21d061e520e8-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Masanori Kondo", Username: "onetap13", Email: "void.kondo@1pac.vn", Password: "onetap", Address: "98 Wall St", PhoneNumber: "0909 999 010", Avatar: "https://ca.slack-edge.com/T0X7R6A5N-U0X7SHF47-8121810751f9-512", Type: models.Business, CompanyID: models.ConvertPointerUint(1), Role: models.UserCompanyMember},
		{Name: "Unknown 1", Username: "onetap14", Email: "dummy1@onetap.vn", Password: "onetap", Address: "100 Wall St", PhoneNumber: "0909 999 011", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Business, CompanyID: models.ConvertPointerUint(2), Role: models.UserCompanyManager},
		{Name: "Unknown 2", Username: "onetap15", Email: "dummy2@onetap.vn", Password: "onetap", Address: "102 Wall St", PhoneNumber: "0909 999 012", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
		{Name: "Unknown 3", Username: "onetap16", Email: "dummy3@onetap.vn", Password: "onetap", Address: "104 Wall St", PhoneNumber: "0909 999 013", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
		{Name: "Unknown 4", Username: "onetap17", Email: "dummy4@onetap.vn", Password: "onetap", Address: "106 Wall St", PhoneNumber: "0909 999 014", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
		{Name: "Unknown 5", Username: "onetap18", Email: "dummy5@onetap.vn", Password: "onetap", Address: "108 Wall St", PhoneNumber: "0909 999 015", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
		{Name: "Unknown 6", Username: "onetap19", Email: "dummy6@onetap.vn", Password: "onetap", Address: "110 Wall St", PhoneNumber: "0909 999 016", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
		{Name: "Unknown 7", Username: "onetap20", Email: "dummy7@onetap.vn", Password: "onetap", Address: "112 Wall St", PhoneNumber: "0909 999 017", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
		{Name: "Unknown 8", Username: "onetap21", Email: "dummy8@onetap.vn", Password: "onetap", Address: "114 Wall St", PhoneNumber: "0909 999 018", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
		{Name: "Unknown 9", Username: "onetap22", Email: "dummy9@onetap.vn", Password: "onetap", Address: "116 Wall St", PhoneNumber: "0909 999 019", Avatar: fmt.Sprintf("%s%s", os.Getenv("app_host"), DefaultAvatar), Type: models.Personal, Role: models.UserStandard},
	}

	for _, user := range users {
		if err := repoUser.Create(&user); err != nil {
			seed.log.Error("failed to load seed: User", zap.Error(err))
		}
	}
}
