package env

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type IENV struct {
	APP_NAME            string
	APP_PORT            string
	APP_PUBLIC          string
	APP_ACTIVE_LOGGING  string
	APP_DB_BUILD_BACKUP string
	APP_TRUSTED_PROXIES string
	APP_SECRET_KEY      string
	DB_CONNECTION       string
	DB_HOST             string
	DB_PORT             string
	DB_DATABASE         string
	DB_USERNAME         string
	DB_PASSWORD         string
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
	ENV.APP_PUBLIC = os.Getenv("APP_PUBLIC")
	ENV.APP_ACTIVE_LOGGING = os.Getenv("APP_ACTIVE_LOGGING")
	ENV.APP_DB_BUILD_BACKUP = os.Getenv("APP_DB_BUILD_BACKUP")
	ENV.APP_TRUSTED_PROXIES = os.Getenv("APP_TRUSTED_PROXIES")
	ENV.APP_SECRET_KEY = os.Getenv("APP_SECRET_KEY")
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

func (ienv *IENV) GetStringArray(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.Split(value, "|")
}

func (ienv *IENV) GetBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return ienv.ParseBool(value)
}

func (ienv *IENV) GetInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return ienv.Atoi(value)
}

func (ienv *IENV) GetFloat64(key string, defaultValue float64) float64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return ienv.ParseFloat64(value)
}

func (ienv *IENV) ParseBool(value string) bool {
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false // nilai default
	}
	return result
}

func (ienv *IENV) Atoi(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0 // nilai default
	}
	return result
}

func (ienv *IENV) ParseFloat64(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0 // nilai default
	}
	return result
}
