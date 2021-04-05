package repositories

import (
	"testing"

	"github.com/michaelt0520/nfc-card/config"
	"github.com/stretchr/testify/assert"
)

type ConfigTesting struct {
	conf  *config.Config
	uRepo *UserRepo
}

var configTesting *ConfigTesting

func TestDAOInit(t *testing.T) {
	assert := assert.New(t)

	// load config
	conf := config.GetConfig()

	// init repositories
	err := InitDAO(conf)
	assert.Nil(err)

	configTesting = &ConfigTesting{
		conf:  conf,
		uRepo: NewUserRepo(),
	}
}
