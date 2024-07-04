package config

import (
	"simple-api/boostrap/module/env"
)

type AppConfig struct {
	Config
	Value *IAppConfig
}

type IAppConfig struct {
	APP_NAME            string
	APP_PORT            string
	APP_PUBLIC          string
	APP_ACTIVE_LOGGING  bool
	APP_DB_BUILD_BACKUP bool
	APP_TRUSTED_PROXIES []string
	APP_SECRET_KEY      string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (config *AppConfig) Load() *IAppConfig {

	// Check if Config is already initialized
	if config.Value != nil {
		return config.Value
	}

	// Initialize ENV
	config.Value = &IAppConfig{}

	env := env.Load()

	config.Value = &IAppConfig{
		APP_NAME:            env.Get("APP_NAME", "main"),
		APP_PORT:            env.Get("APP_PORT", ":8080"),
		APP_PUBLIC:          env.Get("APP_PUBLIC", "public"),
		APP_ACTIVE_LOGGING:  env.GetBool("APP_ACTIVE_LOGGING", true),
		APP_DB_BUILD_BACKUP: env.GetBool("APP_DB_BUILD_BACKUP", true),
		APP_TRUSTED_PROXIES: env.GetStringArray("APP_TRUSTED_PROXIES", []string{"0.0.0.0/0"}),
		APP_SECRET_KEY:      env.Get("APP_SECRET_KEY", "123456789"),
	}

	return config.Value
}
