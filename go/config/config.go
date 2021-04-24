package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
		db, err := gorm.Open("sqlite3", os.Getenv("DB_PATH"))
		return db, err
	case "postgres":
		connect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
		db, err := gorm.Open("postgres", connect)
		return db, err
	default:
		err := fmt.Errorf("invalid variable 'DB_TYPE'=%s", dbType)
		return &gorm.DB{}, err
	}
}
