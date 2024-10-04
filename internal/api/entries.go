package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/utils"
	_ "github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

type PageLoadData struct {
	EntryData any
	Logbookid int
	Limit     int
	Offset    int
	LoadMore  bool
}

func GetEntriesFromLogbook(w http.ResponseWriter, r *http.Request) {
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

	// Instantiate entry query parameters
	queryParams := database.GetEntriesFromLogbookParams{
		Logbookid: int32(logbookId),
	}
	var offset int
	var limit int

	// Check for specific limit values or use defaults
	if limit_str == "" {
		queryParams.Limit = 10
	} else {
		offset, err = strconv.Atoi(offset_str)
		if err != nil {
			log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	// Check for specific offset values or use defaults
	if offset_str == "" {
		queryParams.Offset = 0
	} else {
		limit, err = strconv.Atoi(limit_str)
		if err != nil {
			log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Load read values to query parameters
	queryParams.Limit = int32(limit)
	queryParams.Offset = int32(offset)

	// Run query
	data, err := database.New(AppData.Conn).GetEntriesFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Increment offset count on request for loading older entries
	// If the return row count matches the limit value, then allow to load more
	var page_data PageLoadData
	if len(data) == limit {
		page_data = PageLoadData{data, logbookId, limit, limit + offset, true}
		log.Info().Msg("Loading scroll trigger...")
	} else {
		page_data = PageLoadData{data, logbookId, limit, limit + offset, false}
		log.Info().Msg("No more entries to load...")
	}

	// Render the html
	utils.RenderTemplate(AppData, w, "com-logbook-entry", page_data)

}

func GetEntryFromLogbook(w http.ResponseWriter, r *http.Request) {
	logbookId, entryId, err := utils.ExtractIdsFromRoute(r)
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams := database.GetEntryFromLogbookParams{
		Entryid:   int32(entryId),
		Logbookid: int32(logbookId),
	}
	data, err := database.New(AppData.Conn).GetEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	utils.RenderTemplate(AppData, w, "com-logbook-entry", data)
}

func InsertNewEntryInLogbook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var queryParams database.InsertNewEntryInLogbookParams
	queryParams.Title = title
	queryParams.Description = description
	log.Info().Msg(title + ": " + description)
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queryParams.Createdby = 1 // FIXME: This needs to reflect who created the entry
	queryParams.Logbookid = int32(logbookId)
	_, err = database.New(AppData.Conn).InsertNewEntryInLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

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
	result, err := database.New(AppData.Conn).DeleteEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rows_affected, err := result.RowsAffected()
	if rows_affected > 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func ModalCreate(w http.ResponseWriter, r *http.Request) {
	logbookId, err := strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := database.Entry{Logbookid: int32(logbookId)}
	utils.RenderTemplate(AppData, w, "modal", data)
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

	data, err := database.New(AppData.Conn).GetEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}

	utils.RenderTemplate(AppData, w, "modal", data)
}

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
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var queryParams database.UpdateEntryFromLogbookParams
	queryParams.Title = title
	queryParams.Description = description
	queryParams.Logbookid = int32(logbookId)
	queryParams.Entryid = int32(entryId)

	_, err = database.New(AppData.Conn).UpdateEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
