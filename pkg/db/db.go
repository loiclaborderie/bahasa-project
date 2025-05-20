package db

import (
	"fmt"
	"log"

	"github.com/loiclaborderie/bahasa-project/config"
	"github.com/loiclaborderie/bahasa-project/internal/user"
	"github.com/loiclaborderie/bahasa-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	dbConfig, err := config.GetDatabaseConfig()

	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Db_name, dbConfig.Ssl_mode)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	db.AutoMigrate(&user.User{}, &models.Module{}, &models.Dialogue{}, &models.DialogueLine{}, &models.VocabularyItem{}, &models.VocabularyList{})

	migrateModels(db)

	DB = db
	return db
}

func migrateModels(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&user.User{},
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

func GetDB() *gorm.DB {
	return DB
}
