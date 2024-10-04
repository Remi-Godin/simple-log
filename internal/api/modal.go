package api

import (
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func ModalCreate(w http.ResponseWriter, r *http.Request) {
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := database.Entry{Logbookid: int32(logbookId)}
	utils.RenderTemplate(global.AppData, w, "modal", data)
}

func ModalEdit(w http.ResponseWriter, r *http.Request) {
	logbookId, entryId, err := utils.ExtractIdsFromRoute(r)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var queryParams database.GetEntryFromLogbookParams
	queryParams.Entryid = int32(entryId)
	queryParams.Logbookid = int32(logbookId)

	data, err := database.New(global.AppData.Conn).GetEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}

	utils.RenderTemplate(global.AppData, w, "modal", data)
}
