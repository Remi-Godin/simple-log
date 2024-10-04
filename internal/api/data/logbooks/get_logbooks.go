package logbooks

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/auth"
	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	_ "github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

type MultipleLogbookData struct {
	Logbooks []database.GetLogbooksOwnedByRow
	Links    map[string]string
}

func newMultipleLogbookData() MultipleLogbookData {
	return MultipleLogbookData{
		Links: make(map[string]string),
	}
}

func GetLogbooksOwnedBy(w http.ResponseWriter, r *http.Request) {
	// Parse logbook ID from URL
	// TODO: Get email from JWT user
	email, err := auth.ExtractUserFromJwt(r)
	if err != nil {
		log.Error().Err(err).Msg("Could not get user information")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestParams := r.URL.Query()
	limitStr := requestParams.Get("limit")
	offsetStr := requestParams.Get("offset")

	// Get offset and limit from request
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

	queryParams := database.GetLogbooksOwnedByParams{
		Ownedby: email,
		Limit:   int32(limit),
		Offset:  int32(offset),
	}

	logbooks, err := database.New(global.AppData.Conn).GetLogbooksOwnedBy(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Query failure")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := newMultipleLogbookData()
	data.Logbooks = logbooks
	if len(logbooks) == limit {
		data.Links["LoadMore"] = fmt.Sprintf("/logbooks?limit=%d&offset=%d", limit, offset+limit)
	}

	utils.RenderTemplate(global.AppData, w, "com-logbook", data)
}
