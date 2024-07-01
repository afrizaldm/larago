package config

import "simple-api/cmd/env"

type IConfig struct {
	APP_NAME   string
	APP_PORT   string
	APP_PUBLIC string
}

var Config *IConfig = nil

func Load() *IConfig {

	// Check if Config is already initialized
	if Config != nil {
		return Config
	}

	// Initialize ENV
	Config = &IConfig{}

	env := env.Load()

	Config = &IConfig{
		APP_NAME:   env.Get("APP_NAME", "main"),
		APP_PORT:   env.Get("APP_PORT", ":8080"),
		APP_PUBLIC: env.Get("APP_PUBLIC", "public"),
	}

	return Config
}
