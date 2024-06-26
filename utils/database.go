package utils

import (
	"fmt"
	"kgpsc-backend/models"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {
	user_mig_err := db.AutoMigrate(&models.User{})
	if user_mig_err != nil {
		log.Err(user_mig_err).Msg("User migration error.")
		return user_mig_err
	}

	return nil
}

func GetDB() (db *gorm.DB, err error) {

	DB_USERNAME := os.Getenv("DATABASE_USERNAME")
	DB_PASSWORD := os.Getenv("DATABASE_PASSWORD")
	DB_NAME := os.Getenv("DATABASE_NAME")
	DB_HOST := os.Getenv("DATABASE_HOST")
	DB_PORT := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DB_HOST,
		DB_USERNAME,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Err(err).Msg("Database open error.")

		return nil, err
	}

	return db, nil
}
