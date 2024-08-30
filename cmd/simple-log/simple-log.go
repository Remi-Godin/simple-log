package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Remi-Godin/simple-log/internal/database"
	_ "github.com/Remi-Godin/simple-log/internal/database"
)

type Env struct {
	postgres_user     string
	postgres_password string
	postgres_db       string
	db_addr           string
	db_port           string
	port              string
}

var conn *sql.DB

func loadEnvVars() *Env {
	env := Env{}

	env.postgres_user = os.Getenv("POSTGRES_USER")
	env.postgres_password = os.Getenv("POSTGRES_PASSWORD")
	env.postgres_db = os.Getenv("POSTGRES_DB")
	env.db_addr = os.Getenv("DB_ADDR")
	env.db_port = os.Getenv("DB_PORT")
    env.port = os.Getenv("PORT")

	return &env
}

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
	env := loadEnvVars()
	log.Info().Msg("env variables succesffuly read!")

	// Connect to database
	log.Info().Msg("Initiating database connection...")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.db_addr, env.db_port, env.postgres_user, env.postgres_password, env.postgres_db)
	conn, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not establish database connection")
	}
	defer conn.Close()
	log.Info().Msg("Database connection successful!")

	// Start new mux
	mux := http.NewServeMux()

	// Set Handlers
	//mux.HandleFunc("GET api/v1/logbook/{logbookId}/entries", GetAllEntriesFromLogbook)
    mux.HandleFunc("GET /", index)
	mux.HandleFunc("GET /api/v1/logbook/{logbookId}/entries", GetEntriesFromLogbook)
	mux.HandleFunc("GET /api/v1/logbook/{logbookId}/entries/{entryId}", GetEntryFromLogbook)


	// Start server
    err = http.ListenAndServe(env.db_addr+":"+env.port, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failure")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
    log.Info().Msg("Yup, this is the index")
}


func GetEntriesFromLogbook(w http.ResponseWriter, r *http.Request) {
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
        w.WriteHeader(http.StatusBadRequest)
        return
	}
    request_params := r.URL.Query()
    limit_str := request_params.Get("limit")
    offset_str := request_params.Get("offset")
    if limit_str == "" || offset_str == "" {
        result,err := database.New(conn).GetAllEntriesFromLogbook(r.Context(), int32(logbookId))
        if err != nil {
		    log.Error().Err(err).Msg("Could not complete database query")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        enc := json.NewEncoder(w)
        enc.Encode(result)
        return
    }
    offset,err := strconv.Atoi(offset_str)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
        w.WriteHeader(http.StatusBadRequest)
        return
	}
    limit,err := strconv.Atoi(limit_str)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
        w.WriteHeader(http.StatusBadRequest)
        return
	}
    query_params := database.GetEntriesFromLogbookParams{
        Logbookid: int32(logbookId),
        Offset: int32(offset),
        Limit: int32(limit),
    }
    result,err := database.New(conn).GetEntriesFromLogbook(r.Context(),query_params)
    if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    enc := json.NewEncoder(w)
    enc.Encode(result)
    return

}

func GetEntryFromLogbook(w http.ResponseWriter, r *http.Request) {

}

func GetLogbookOwnerFromEntry(w http.ResponseWriter, r *http.Request) {

}

func InsertNewEntryInLogbook(w http.ResponseWriter, r *http.Request) {

}

func DeleteEntryFromLogbook(w http.ResponseWriter, r *http.Request) {

}

func GetLogbooksOwnedBy(w http.ResponseWriter, r *http.Request) {

}

func InsertNewLogbook(w http.ResponseWriter, r *http.Request) {

}

func DeleteLogbook(w http.ResponseWriter, r *http.Request) {

}
