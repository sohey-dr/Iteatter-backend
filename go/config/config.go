package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

}

func OpenDB() (*gorm.DB, error) {
	loadEnv()
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
		return db, err
	case "postgres":
		connect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
		db, err := gorm.Open(postgres.Open(connect), &gorm.Config{})
		return db, err
	default:
		err := fmt.Errorf("invalid variable 'DB_TYPE'=%s", dbType)
		return &gorm.DB{}, err
	}
}
