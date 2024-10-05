package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

type LogbookPageData struct {
	LogbookId    int
	LogbookTitle string
	Links        map[string]string
}

func newLogbookPageData() LogbookPageData {
	return LogbookPageData{
		0,
		"",
		make(map[string]string),
	}
}

func GetLogbook(w http.ResponseWriter, r *http.Request) {
	// Parse URL
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Query database
	logbookData, err := database.New(global.AppData.Conn).GetLogbookData(r.Context(), int32(logbookId))
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Render template
	data := newLogbookPageData()
	data.LogbookId = logbookId
	data.LogbookTitle = logbookData.Title
	data.Links["Entries"] = fmt.Sprintf("/logbook/%d/entries?limit=%d&offset=%d", logbookId, 5, 0)
	utils.RenderTemplate(global.AppData, w, "logbook", data)
}
