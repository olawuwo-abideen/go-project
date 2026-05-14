package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}

// package initializers

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"github.com/olawuwo-abideen/go-project/go-auth/models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectToDb() {

// 	err := godotenv.Load()

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	dsn := os.Getenv("DATABASE_URL")

// 	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal("Failed to connect to database")
// 	}

// 	DB = database

// 	log.Println("Database connected successfully")
// }

// func SyncDatabase() {
// 	DB.AutoMigrate(&models.User{})
// }
