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
