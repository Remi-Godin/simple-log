package logbooks

import (
	"encoding/json"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/rs/zerolog/log"
)

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
	queryParams.Ownedby = "regodin@proton.me"

	// Execute the query
	_, err = database.New(global.AppData.Conn).InsertNewLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}
