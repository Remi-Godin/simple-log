package api

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/Remi-Godin/simple-log/internal/utils/validation"
	"github.com/Remi-Godin/simple-log/internal/utils/validation/validators"
	"github.com/rs/zerolog/log"
)

func ValidateLastName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("There was an error when parsing the form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := validators.NewNameValidator()
	data.FieldName = "last-name"
	data.FieldId = "last-name"
	data.FieldNameText = "Last name"
	data.Links["ValidateField"] = fmt.Sprintf("/validate/last-name")
	data.FieldValue = r.FormValue("last-name")
	err = validation.Validate(r.Context(), data)
	if data.FieldValue != "" {
		if err != nil {
			data.Err = err.Error()
			data.Invalid = false
		} else {
			data.Valid = true
		}
	}

	utils.RenderTemplate(global.AppData, w, "validated-input-field", data)
}
