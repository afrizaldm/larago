package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IENV struct {
	APP_NAME      string
	APP_PORT      string
	DB_CONNECTION string
	DB_HOST       string
	DB_PORT       string
	DB_DATABASE   string
	DB_USERNAME   string
	DB_PASSWORD   string
}

var ENV *IENV = nil

func Load() *IENV {

	// Check if ENV is already initialized
	if ENV != nil {
		return ENV
	}

	// Initialize ENV
	ENV = &IENV{}

	// Memuat variabel lingkungan dari file .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
	}

	// Mengakses variabel lingkungan yang telah dimuat
	ENV.APP_NAME = os.Getenv("APP_NAME")
	ENV.APP_PORT = os.Getenv("APP_PORT")
	ENV.DB_CONNECTION = os.Getenv("DB_CONNECTION")
	ENV.DB_HOST = os.Getenv("DB_HOST")
	ENV.DB_PORT = os.Getenv("DB_PORT")
	ENV.DB_DATABASE = os.Getenv("DB_DATABASE")
	ENV.DB_USERNAME = os.Getenv("DB_USERNAME")
	ENV.DB_PASSWORD = os.Getenv("DB_PASSWORD")

	return ENV
}

func (ienv *IENV) Get(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
