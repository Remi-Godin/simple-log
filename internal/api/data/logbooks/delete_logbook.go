package logbooks

import (
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/rs/zerolog/log"
)

func DeleteLogbook(w http.ResponseWriter, r *http.Request) {
	// Get logbook id from url
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Execute query
	result, err := database.New(global.AppData.Conn).DeleteLogbook(r.Context(), int32(logbookId))
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
	// if deleted, return 200
	rowsAffected, err := result.RowsAffected()
	if rowsAffected > 0 {
		w.Header().Add("HX-Redirect", "/logbook")
		return
	}
	// if nothing got deleted, then no content
	w.WriteHeader(http.StatusNoContent)
}
