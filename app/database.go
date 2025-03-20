package app

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Username string
	Password string
	DBName   string
	Port     string
	Host     string
}

func ConnectDB(config DBConfig) (*gorm.DB, error) {
	dsn := "user=" + config.Username +
		" password=" + config.Password +
		" dbname=" + config.DBName +
		" port=" + config.Port +
		" host=" + config.Host +
		" sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, errors.New("failed to connect to app")
	}

	return db, nil
}
