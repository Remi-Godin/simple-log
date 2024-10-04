package entries

import (
	"database/sql"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func GetEntryFromLogbook(w http.ResponseWriter, r *http.Request) {
	logbookId, entryId, err := utils.ExtractIdsFromRoute(r)
	if err != nil {
		log.Error().Err(err).Msg("Could not extract IDs")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams := database.GetEntryFromLogbookParams{
		Entryid:   int32(entryId),
		Logbookid: int32(logbookId),
	}
	data, err := database.New(global.AppData.Conn).GetEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	utils.RenderTemplate(global.AppData, w, "com-logbook-entry", data)
}
