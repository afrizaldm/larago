package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IENV struct {
	APP_NAME      string
	DB_CONNECTION string
	DB_HOST       string
	DB_PORT       string
	DB_DATABASE   string
	DB_USERNAME   string
	DB_PASSWORD   string
}

var ENV IENV

func Load() IENV {
	// Memuat variabel lingkungan dari file .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
	}

	// Mengakses variabel lingkungan yang telah dimuat
	ENV.APP_NAME = os.Getenv("APP_NAME")
	ENV.DB_CONNECTION = os.Getenv("DB_CONNECTION")
	ENV.DB_HOST = os.Getenv("DB_HOST")
	ENV.DB_PORT = os.Getenv("DB_PORT")
	ENV.DB_DATABASE = os.Getenv("DB_DATABASE")
	ENV.DB_USERNAME = os.Getenv("DB_USERNAME")
	ENV.DB_PASSWORD = os.Getenv("DB_PASSWORD")

	return ENV
}
