package entries

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func UpdateEntryFromLogbook(w http.ResponseWriter, r *http.Request) {
	logbookId, entryId, err := utils.ExtractIdsFromRoute(r)
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
		log.Error().Err(err).Msg("Could not parse entry data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var queryParams database.UpdateEntryFromLogbookParams
	queryParams.Title = title
	queryParams.Description = description
	queryParams.Logbookid = int32(logbookId)
	queryParams.Entryid = int32(entryId)

	_, err = database.New(global.AppData.Conn).UpdateEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Add("HX-Redirect", fmt.Sprintf("/logbook/%d", logbookId))
}
