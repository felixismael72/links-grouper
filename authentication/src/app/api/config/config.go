package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func ByEnvType(envType string) API {
	log.Info().Msgf("Setting up application for %s environment...", envType)

	if err := loadEnvFile(envType); err != nil {
		log.Fatal().Msgf("Failed to load environment file: %s", err.Error())
	}

	return getConfig()
}

func loadEnvFile(envType string) error {
	switch envType {
	case "production":
		return godotenv.Load(".env.production")
	case "development", "":
		return godotenv.Load(".env.development")
	default:
		return errors.New("unsupported environment type")
	}
}

func getConfig() API {
	var config API

	config.Server.Host = os.Getenv("API_HOST")
	config.Server.Port, _ = strconv.Atoi(os.Getenv("API_PORT"))
	config.PG.Username = os.Getenv("PG_USER")
	config.PG.Password = os.Getenv("PG_PASSWORD")
	config.PG.Server.Host = os.Getenv("PG_HOST")
	config.PG.Server.Port, _ = strconv.Atoi(os.Getenv("PG_PORT"))
	config.PG.DBName = os.Getenv("PG_DB_NAME")
	config.JWT.Secret = os.Getenv("SERVER_SECRET")

	return config
}
