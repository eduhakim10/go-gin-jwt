package initializers

import (
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var name = "Edu Arif Rahman Hakim"
var DB *gorm.DB

func NowDB() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Get database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbTimezone := os.Getenv("DB_TIMEZONE")

	// Construct DSN
	dsn := "user=" + dbUser +
		" password=" + dbPassword +
		" dbname=" + dbName +
		" port=" + dbPort +
		" host=" + dbHost +
		" sslmode=disable TimeZone=" + dbTimezone

		database, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
	if err != nil {

		panic(err)

	} else {
		fmt.Println("connect cuy")
	}
	DB = database
	//database.AutoMigrate(&models.User)

}
