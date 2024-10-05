package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	_ "github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

type EntriesData struct {
	EntryData []database.Entry
	Links     map[string]string
}

func newEntriesData() EntriesData {
	return EntriesData{
		Links: make(map[string]string),
	}
}

func GetEntriesFromLogbook(w http.ResponseWriter, r *http.Request) {
	// Parse logbook ID from URL
	log.Info().Msg("Getting entries from logbook")
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse URL parameters
	request_params := r.URL.Query()
	limit_str := request_params.Get("limit")
	offset_str := request_params.Get("offset")

	// Get offset and limit from request
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

	// Set query parameters
	queryParams := database.GetEntriesFromLogbookParams{
		Logbookid: int32(logbookId),
		Limit:     int32(limit),
		Offset:    int32(offset),
	}

	// Run query
	entries, err := database.New(global.AppData.Conn).GetEntriesFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := newEntriesData()
	data.EntryData = entries
	if len(entries) == limit {
		data.Links["Entries"] = fmt.Sprintf("/logbook/%d/entries?limit=%d&offset=%d", logbookId, limit, offset+limit)
	}

	// Render the html
	utils.RenderTemplate(global.AppData, w, "com-logbook-entry", data)

}
