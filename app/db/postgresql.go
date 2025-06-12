package db

import (
	"database/sql"
	"fiber-starter/app/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var sqlDB *sql.DB

func ConnectDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	err = db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	sqlDBInstance, err := db.DB()
	if err != nil {
		log.Fatal("Error getting database instance: ", err)
	}

	sqlDB = sqlDBInstance
	DB = db

	log.Println("Connected to database")
}

func CloseDB() {
	if sqlDB != nil {
		log.Println("Closing database connection")
		if err := sqlDB.Close(); err != nil {
			log.Println("Error closing database connection: ", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
