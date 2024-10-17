package entries

import (
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/rs/zerolog/log"
)

func InsertNewEntryInLogbook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Could not parse new entry form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		log.Error().Err(err).Msg("Could not parse title and description")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var queryParams database.InsertNewEntryInLogbookParams
	queryParams.Title = title
	queryParams.Description = description
	log.Info().Msg(title + ": " + description)
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Could not get logbook ID from path")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams.Createdby = "regodin@proton.me" // FIXME: This needs to reflect who created the entry
	queryParams.Logbookid = int32(logbookId)
	_, err = database.New(global.AppData.Conn).InsertNewEntryInLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
