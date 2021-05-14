package db

import (
	"fmt"

	"github.com/michaelt0520/nfc-card/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// InitRepository ...
func ConnectDatabase(conf *config.Config) error {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DBUserName, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)

	fmt.Println(connectStr)
	connectedDb, err := gorm.Open(mysql.Open(connectStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	db = connectedDb

	return nil
}

func Database() *gorm.DB {
	return db
}
