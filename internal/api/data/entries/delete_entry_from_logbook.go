package entries

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func DeleteEntryFromLogbook(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Starting deletion")
	logbookId, entryId, err := utils.ExtractIdsFromRoute(r)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams := database.DeleteEntryFromLogbookParams{
		Entryid:   int32(entryId),
		Logbookid: int32(logbookId),
	}
	result, err := database.New(global.AppData.Conn).DeleteEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected > 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
