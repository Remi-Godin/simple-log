package logbooks

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/auth"
	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/rs/zerolog/log"
)

func InsertLogbook(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Inserting new logbook")
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Could not parse logbook form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := auth.ExtractUserFromJwt(r)
	if err != nil {
		log.Error().Err(err).Msg("Could not extract user from JWT")
	}

	queryParams := database.InsertNewLogbookParams{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Ownedby:     user,
	}
	_, err = database.New(global.AppData.Conn).InsertNewLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert new logbook")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Redirect", "/logbook")
	w.WriteHeader(http.StatusCreated)
}
