package main

import (
	"fmt"
	"os"
	"path/filepath"

	"authentication/src/app/api/config"
	"authentication/src/app/api/endpoints/routes"
	"authentication/src/app/api/flags"
	"authentication/src/core/domain/authorization"
	"authentication/src/infra/postgres"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	validateExecutionPath()

	flags := flags.Parse()
	apiConfig := config.ByEnvType(*flags.EnvType)
	setupPG(apiConfig.PG)
	setupJWT(apiConfig.JWT)
	setupServer(apiConfig.Server)
}

func setupPG(config config.PG) {
	postgres.SetupCredentials(
		config.Server.Host,
		config.Server.Port,
		config.Username,
		config.Password,
		config.DBName,
	)
}

func setupJWT(config config.JWT) {
	authorization.SecretKey = config.Secret
}

func setupServer(config config.Server) {
	router := echo.New()
	routes.Load(router)
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	router.Logger.Fatal(router.Start(address))
}

func validateExecutionPath() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal().Msg("Failed to get the current working directory")
	}

	filePath := filepath.Join(dir, "go.mod")
	if _, err := os.Stat(filePath); err != nil {
		log.Fatal().Msg("Program must run from the project's root.")
	}
}
