package repositories

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/michaelt0520/nfc-card/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// InitRepository ...
func InitRepository(conf *config.Config) (err error) {
	connectStr := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBUserName, conf.DBPassword, conf.DBPort, conf.DBName)
	connectedDb, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	db = connectedDb

	return nil
}

// GetDB : getter
func GetDB() *gorm.DB {
	return db
}

type Repository interface {
	All(result interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error)
	Find(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error)
	Where(result interface{}, data map[string]interface{}, scopes ...func(db *gorm.DB) *gorm.DB) (*gorm.DB, error)
	Create(model interface{}) (*gorm.DB, error)
	Update(model interface{}, data map[string]interface{}) (*gorm.DB, error)
	Destroy(model interface{}) (*gorm.DB, error)
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("per_page"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
