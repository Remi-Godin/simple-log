package forms

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func NewEntry(w http.ResponseWriter, r *http.Request) {
	logbookId := r.PathValue("logbookId")
	log.Warn().Msg(logbookId)

	data := NewFormData("Create Entry")
	data.FormDesc = "Enter entry details."

	data.FormSubmissionLink = fmt.Sprintf("/logbook/%s/entries", logbookId)
	data.FormFields = append(data.FormFields, "/field/title")
	data.FormFields = append(data.FormFields, "/field/description")

	utils.RenderTemplate(global.AppData, w, "com-modal", data)
}

func EditEntry(w http.ResponseWriter, r *http.Request) {
	logbookId, entryId, err := utils.ExtractIdsFromRoute(r)
	if err != nil {
		log.Error().Err(err).Msg("Could not extract data from path.")
	}

	data := NewFormData("Edit Entry")
	data.FormDesc = ""

	queryParams := database.GetEntryFromLogbookParams{
		Logbookid: int32(logbookId),
		Entryid:   int32(entryId),
	}
	queryResult, err := database.New(global.AppData.Conn).GetEntryFromLogbook(r.Context(), queryParams)
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data.FormSubmissionLink = fmt.Sprintf("/logbook/%d/entries/%d", logbookId, entryId)
	data.ResourceDeletionLink = fmt.Sprintf("/logbook/%d/entries/%d", logbookId, entryId)
	data.Patch = true
	data.FormFields = append(data.FormFields, fmt.Sprintf("/field/title?value=%s", url.QueryEscape(queryResult.Title)))
	data.FormFields = append(data.FormFields, fmt.Sprintf("/field/description?value=%s", url.QueryEscape(queryResult.Description)))

	utils.RenderTemplate(global.AppData, w, "com-modal", data)
}
