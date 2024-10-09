package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Remi-Godin/simple-log/internal/api"
	_ "github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
)

func main() {
	// Set logger for pretty printing to console
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Server starting up...")

	// Load env file
	log.Info().Msg("Loading env file...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not load .env file")
	}
	log.Info().Msg("env file succesffuly loaded!")

	// Create env variables struct
	log.Info().Msg("Reading env variables...")
	global.AppData = utils.AppData{}
	global.AppData.Env = *utils.LoadEnvVars()

	// Connect to database
	log.Info().Msg("Initiating database connection...")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		global.AppData.Env.Db_addr, global.AppData.Env.Db_port, global.AppData.Env.Postgres_user, global.AppData.Env.Postgres_password, global.AppData.Env.Postgres_db)
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not establish database connection")
	}
	defer conn.Close()
	log.Info().Msg("Database connection successful!")
	global.AppData.Conn = conn
	log.Err(conn.Ping())

	// Parse templates
	global.AppData.Tmpl = template.Must(template.ParseGlob("./web/templates/*.html"))

	// Start new mux
	mux := http.NewServeMux()
	api.SetRoutes(mux)

	// Start server
	log.Info().Msg("Starting server at: " + global.AppData.Env.Db_addr + ":" + global.AppData.Env.Port)
	err = http.ListenAndServe(global.AppData.Env.Db_addr+":"+global.AppData.Env.Port, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failure")
	}
}
