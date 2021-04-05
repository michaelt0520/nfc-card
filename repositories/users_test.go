package repositories

import (
	"testing"

	"github.com/michaelt0520/nfc-card/models"
	"github.com/stretchr/testify/assert"
)

type UserDAODataTesting struct {
	userNameTestingData string
	userID              uint32
}

var udt *UserDAODataTesting

func UserDAOClearDataTesting() error {
	err := GetDB().Unscoped().Where("username like ?", "%"+udt.userNameTestingData+"%").Delete(&models.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func TestUserDAO(t *testing.T) {
	udt = &UserDAODataTesting{
		userNameTestingData: "wakuwaku-user-name-testing-prefix-",
	}
}

func TestUserDAOFind(t *testing.T) {
	assert := assert.New(t)

	// clear data before test
	err := UserDAOClearDataTesting()
	assert.Nil(err)

	// init data-testing
	TestUserDAO(t)

	// add new user
	userModel := models.User{
		Username:        udt.userNameTestingData,
		Email:           udt.userNameTestingData + "@gmail.com",
		Name:            "testing-waku",
		LastSignInIP:    "::1",
		CurrentSignInIP: "::1",
	}
	err = GetDB().Save(&userModel).Error
	assert.Nil(err)

	udt.userID = userModel.ID

	// Find
	res, err := configTesting.ud.Find(int(udt.userID))
	assert.Nil(err)
	assert.NotNil(res)
	assert.Equal(udt.userNameTestingData, res.Username, "user_name is not equal")
	assert.Equal(udt.userNameTestingData+"@gmail.com", res.Email, "email is not equal")
	assert.Equal("testing-waku", res.Name, "name is not equal")
}

func TestUserDAOFinish(t *testing.T) {
	assert := assert.New(t)

	// clear data before test
	err := UserDAOClearDataTesting()
	assert.Nil(err)
}
