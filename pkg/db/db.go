package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/loiclaborderie/bahasa-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "bahasa_project")
	sslmode := getEnv("DB_SSLMODE", "disable")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Module{}, &models.Dialogue{}, &models.DialogueLine{}, &models.VocabularyItem{}, &models.VocabularyList{})

	log.Println("Connected to database successfully")

	migrateModels(db)

	return db
}

func migrateModels(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Module{},
		&models.Dialogue{},
		&models.DialogueLine{},
		&models.VocabularyItem{},
		&models.VocabularyList{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrations completed successfully")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
