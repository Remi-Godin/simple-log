package entries

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

type MultipleEntryData struct {
	Entries []database.Entry
	Links   map[string]string
}

func newMultipleEntryData() MultipleEntryData {
	return MultipleEntryData{
		Links: make(map[string]string),
	}
}

func GetEntriesFromLogbook(w http.ResponseWriter, r *http.Request) {
	// Parse logbook ID from URL
	// TODO: Get email from JWT user

	logbookIdStr := r.PathValue("logbookId")
	requestParams := r.URL.Query()
	limitStr := requestParams.Get("limit")
	offsetStr := requestParams.Get("offset")

	// Get offset and limit from request
	logbookId, err := strconv.Atoi(logbookIdStr)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryParams := database.GetEntriesFromLogbookParams{
		Logbookid: int32(logbookId),
		Limit:     int32(limit),
		Offset:    int32(offset),
	}

	entries, err := database.New(global.AppData.Conn).GetEntriesFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Query failure")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := newMultipleEntryData()
	data.Entries = entries
	if len(entries) == limit {
		data.Links["LoadMore"] = fmt.Sprintf("/logbook/%d/entries?limit=%d&offset=%d", logbookId, limit, offset+limit)
	}

	utils.RenderTemplate(global.AppData, w, "com-logbook-entry", data)
}
