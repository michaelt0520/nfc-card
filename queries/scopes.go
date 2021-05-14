package queries

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildWhere(_filters map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for k, v := range _filters {
			if v != nil {
				db.Where(k, v)
			} else {
				db.Where(k)
			}
		}

		return db
	}
}

func BuildOr(_filters map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for k, v := range _filters {
			if v != nil {
				db.Or(k, v)
			} else {
				db.Or(k)
			}
		}

		return db
	}
}

func BuildPreload(_associations map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for k, v := range _associations {
			if v != nil {
				db.Preload(k, v)
			} else {
				db.Preload(k)
			}
		}

		return db
	}
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
			pageSize = 100
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
