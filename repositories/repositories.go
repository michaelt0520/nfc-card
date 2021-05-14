package repositories

import (
	"context"

	"github.com/michaelt0520/nfc-card/db"
	"gorm.io/gorm"
)

// GetDB : getter
func GetDB() *gorm.DB {
	query := db.Database().WithContext(context.Background())
	return query
}
