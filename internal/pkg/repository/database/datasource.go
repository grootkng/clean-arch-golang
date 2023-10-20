package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/grootkng/clean-arch-golang/config"
)

func Db() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetEnv().DB), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

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
