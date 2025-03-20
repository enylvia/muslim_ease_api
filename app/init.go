package app

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

// config DB on here
var (
	configDatabase DBConfig
)

type ConfigDB struct {
	DB *gorm.DB
}

func loadEnvConfig() DBConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	configDatabase := DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
	}

	return configDatabase
}
func Init() ConfigDB {
	configDatabase := loadEnvConfig()

	dbDarlink, err := ConnectDB(configDatabase)
	if err != nil {
		log.Fatalf("failed to connect to dbDarlink: %v", err)
	}

	return ConfigDB{
		DB: dbDarlink,
	}
}
