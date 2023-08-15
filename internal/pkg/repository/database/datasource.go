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
