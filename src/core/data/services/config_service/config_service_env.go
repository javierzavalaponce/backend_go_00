package config_service

import (
	"log"
	"os"

	coreModels "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/models"
	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"

	"github.com/joho/godotenv"
)

type ConfigServiceEnv struct {
	config *coreModels.Config
}

var _ core_services.ConfigService = (*ConfigServiceEnv)(nil)

func NewConfigServiceEnv() *ConfigServiceEnv {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	config := &coreModels.Config{
		ApiPort: os.Getenv("API_PORT"),
		ApiHost: os.Getenv("API_HOST"),

		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbSslMode:  os.Getenv("DB_SSL_MODE"),

		LogLevel:      os.Getenv("LOG_LEVEL"),
		LogOutput:     os.Getenv("LOG_OUTPUT"),
		SessionSecret: os.Getenv("SESSION_SECRET"),

		MigrationsPath: os.Getenv("MIGRATIONS_PATH"),
		SeedersPath:    os.Getenv("SEEDERS_PATH"),
	}
	return &ConfigServiceEnv{config: config}
}

func (e *ConfigServiceEnv) Read() *coreModels.Config {
	return e.config
}
