package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"os"

	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("SSLMODE")

	d, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password+" sslmode="+sslmode)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
