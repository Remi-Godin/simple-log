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

func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := validators.NewEmailValidator()
	data.FieldValue = r.FormValue("email")
	data.Links["ValidateField"] = fmt.Sprintf("/validate/email")
	data.FieldValue = r.FormValue("email")
	err = validation.Validate(r.Context(), data)
	if data.FieldValue != "" {
		if err != nil {
			log.Err(err)
			data.Err = err.Error()
			data.Invalid = true
		} else {
			data.Valid = true
		}
	}

	utils.RenderTemplate(global.AppData, w, "validated-input-field", data)
}
