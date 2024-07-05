package config

import "simple-api/boostrap/module/env"

type DatabaseConfig struct {
	Config
	Value *IDatabaseConfig
}

type IDatabaseConfig struct {
	DB_CONNECTION string
	DB_HOST       string
	DB_PORT       string
	DB_DATABASE   string
	DB_USERNAME   string
	DB_PASSWORD   string
}

func DatabaseConfigLoad() *IDatabaseConfig {
	return NewDatabaseConfig().Load()
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{}
}

func (config *DatabaseConfig) Load() *IDatabaseConfig {

	// Check if Config is already initialized
	if config.Value != nil {
		return config.Value
	}

	// Initialize ENV
	config.Value = &IDatabaseConfig{}

	env := env.Load()

	config.Value = &IDatabaseConfig{
		DB_CONNECTION: env.Get("DB_CONNECTION", "sqlite"),
		DB_HOST:       env.Get("DB_HOST", "127.0.0.1"),
		DB_PORT:       env.Get("DB_PORT", "3306"),
		DB_DATABASE:   env.Get("DB_DATABASE", "goat"),
		DB_USERNAME:   env.Get("DB_USERNAME", "root"),
		DB_PASSWORD:   env.Get("DB_PASSWORD", ""),
	}

	return config.Value
}
