package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db_name  string
	Ssl_mode string
}

func GetDatabaseConfig() (DbConfig, error) {

	if err := godotenv.Load(); err != nil {
		return DbConfig{}, errors.New("we could not load any env file")
	}

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "bahasa_project")
	sslmode := getEnv("DB_SSLMODE", "disable")

	return DbConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Db_name:  dbname,
		Ssl_mode: sslmode,
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
