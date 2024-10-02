package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
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
var tmpl *template.Template

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

func renderTemplate(w http.ResponseWriter, tmpl_name string, data any) {
	log.Info().Msg("Rendering template")
	err := tmpl.ExecuteTemplate(w, tmpl_name, data)
	if err != nil {
		log.Error().Err(err).Msg("Could not execute template")
	}
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

	// Parse templates
	tmpl = template.Must(template.ParseGlob("./web/templates/*.html"))

	// Start new mux
	mux := http.NewServeMux()

	// Set Handlers
	//mux.HandleFunc("GET api/v1/logbook/{logbookId}/entries", GetAllEntriesFromLogbook)
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("web/styles"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /logbook", GetLogbooks)
	mux.HandleFunc("GET /logbook/{logbookId}", GetLogbook)
	mux.HandleFunc("GET /logbook/{logbookId}/entries", GetEntriesFromLogbook)
	mux.HandleFunc("GET /logbook/{logbookId}/entries/{entryId}", GetEntryFromLogbook)
	mux.HandleFunc("POST /logbook/{logbookId}/entries", InsertNewEntryInLogbook)
	mux.HandleFunc("POST /logbook", InsertNewLogbook)
	mux.HandleFunc("DELETE /logbook/{logbookId}/entries/{entryId}", DeleteEntryFromLogbook)
	mux.HandleFunc("DELETE /logbook/{logbookId}", DeleteLogbook)
	mux.HandleFunc("GET /modal/create", ModalCreate)
	mux.HandleFunc("GET /modal/edit/{logbookId}/{entryId}", ModalEdit)
	mux.HandleFunc("PATCH /logbook/{logbookId}/entries/{entryId}", UpdateEntryFromLogbook)

	// Start server
	log.Info().Msg("Starting server at: " + env.db_addr + ":" + env.port)
	err = http.ListenAndServe(env.db_addr+":"+env.port, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failure")
	}
}

func UpdateEntryFromLogbook(w http.ResponseWriter, r *http.Request) {
	// Get entry data from request
	// Get owner data from token (not yet implemented)
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entryId, err := strconv.Atoi(r.PathValue("entryId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to parse form but failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var queryParams database.UpdateEntryFromLogbookParams
	queryParams.Title = title
	queryParams.Description = description
	queryParams.Logbookid = int32(logbookId)
	queryParams.Entryid = int32(entryId)

	_, err = database.New(conn).UpdateEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ModalCreate(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "modal", nil)
}

func ModalEdit(w http.ResponseWriter, r *http.Request) {
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entryId, err := strconv.Atoi(r.PathValue("entryId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var queryParams database.GetEntryFromLogbookParams
	queryParams.Entryid = int32(entryId)
	queryParams.Logbookid = int32(logbookId)

	data, err := database.New(conn).GetEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}

	renderTemplate(w, "modal", data)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Yup, this is the index")
}

type PageLoadData struct {
	EntryData any
	Limit     int
	Offset    int
	LoadMore  bool
}

func GetEntriesFromLogbook(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Getting entries from logbook")
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request_params := r.URL.Query()
	limit_str := request_params.Get("limit")
	offset_str := request_params.Get("offset")
	latest_only := request_params.Get("latest_only")
	if limit_str == "" || offset_str == "" {
		data, err := database.New(conn).GetAllEntriesFromLogbook(r.Context(), int32(logbookId))
		if err != nil {
			log.Error().Err(err).Msg("Could not complete database query")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		renderTemplate(w, "com-logbook-entry", data)
		return
	}
	offset, err := strconv.Atoi(offset_str)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(limit_str)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams := database.GetEntriesFromLogbookParams{
		Logbookid: int32(logbookId),
		Offset:    int32(offset),
		Limit:     int32(limit),
	}
	data, err := database.New(conn).GetEntriesFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var page_data PageLoadData
	if latest_only == "true" {
		page_data = PageLoadData{data, 1, 0, false}
	} else if len(data) == limit {
		page_data = PageLoadData{data, limit, limit + offset, true}
	} else {
		page_data = PageLoadData{data, limit, limit + offset, false}
	}
	renderTemplate(w, "com-logbook-entry", page_data)

}

func GetEntryFromLogbook(w http.ResponseWriter, r *http.Request) {
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entryId, err := strconv.Atoi(r.PathValue("entryId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams := database.GetEntryFromLogbookParams{
		Entryid:   int32(entryId),
		Logbookid: int32(logbookId),
	}
	data, err := database.New(conn).GetEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "com-logbook-entry", data)
}

func InsertNewEntryInLogbook(w http.ResponseWriter, r *http.Request) {
	// Get entry data from request
	// Get owner data from token (not yet implemented)
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var queryParams database.InsertNewEntryInLogbookParams
	queryParams.Title = title
	queryParams.Description = description
	log.Info().Msg(title + ": " + description)
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams.Createdby = 1 // FIXME: This needs to reflect who created the entry
	queryParams.Logbookid = int32(logbookId)
	_, err = database.New(conn).InsertNewEntryInLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DeleteEntryFromLogbook(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Starting deletion")
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entryId, err := strconv.Atoi(r.PathValue("entryId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams := database.DeleteEntryFromLogbookParams{
		Entryid:   int32(entryId),
		Logbookid: int32(logbookId),
	}
	result, err := database.New(conn).DeleteEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rows_affected, err := result.RowsAffected()
	if rows_affected > 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func GetLogbooksOwnedBy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func InsertNewLogbook(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Inserting new logbook")
	// Decode json data in body
	decoder := json.NewDecoder(r.Body)
	var queryParams database.InsertNewLogbookParams
	err := decoder.Decode(&queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not decode json payload.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Set Ownedby to the user that sent the request
	// FIXME: This needs to reflect whoever created the logbook
	queryParams.Ownedby = 1

	// Execute the query
	_, err = database.New(conn).InsertNewLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteLogbook(w http.ResponseWriter, r *http.Request) {
	// Get logbook id from url
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Execute query
	result, err := database.New(conn).DeleteLogbook(r.Context(), int32(logbookId))
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
	// if deleted, return 200
	rows_affected, err := result.RowsAffected()
	if rows_affected > 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	// if nothing got deleted, then no content
	w.WriteHeader(http.StatusNoContent)
}

func GetLogbooks(w http.ResponseWriter, r *http.Request) {
	result, err := database.New(conn).GetLogbooksOwnedBy(r.Context(), 1)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}

	enc := json.NewEncoder(w)
	enc.Encode(result)
}

func GetLogbook(w http.ResponseWriter, r *http.Request) {
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := logbookId
	log.Info().Msg(string(data))
	renderTemplate(w, "logbook", nil)
}
