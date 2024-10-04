package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

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
	_, err = database.New(AppData.Conn).InsertNewLogbook(r.Context(), queryParams)
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
	result, err := database.New(AppData.Conn).DeleteLogbook(r.Context(), int32(logbookId))
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
	result, err := database.New(AppData.Conn).GetLogbooksOwnedBy(r.Context(), 1)
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
	utils.RenderTemplate(AppData, w, "logbook", data)
}
