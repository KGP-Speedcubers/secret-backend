package main

import (
	"flag"
	"kgpsc-backend/server"
	"kgpsc-backend/utils"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	envFile := flag.String("envFile", ".env", "A file to load environment variables from.")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msgf("Attempting to load environment variables from %s.", *envFile)
	dotenv_err := godotenv.Load(*envFile)

	if dotenv_err != nil {
		log.Warn().Msgf("Failed to load environment variables from %s.", *envFile)
	} else {
		log.Info().Msgf("Successfully loaded environment variables from %s.", *envFile)
	}

	db, db_err := utils.GetDB()

	if db_err != nil {
		log.Fatal().Err(db_err).Msg("Error connecting to the database.")
	}

	mig_err := utils.MigrateModels(db)
	if mig_err != nil {
		log.Fatal().Err(mig_err).Msg("Database migration error.")
	}

	log.Info().Msg("Creating mux router")
	router := server.NewRouter(db)

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	corsObj := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ORIGINS_ALLOWED")},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	log.Info().Msg("Starting server on port : " + port)
	err := http.ListenAndServe(
		":"+port,
		corsObj.Handler(router),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Error starting server.")
	}
}
