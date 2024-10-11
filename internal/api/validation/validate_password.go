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

func ValidatePasswordStrength(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := validators.NewPasswordValidator()
	data.Links["ValidateField"] = fmt.Sprintf("/validate/password")
	data.FieldValue = r.FormValue("password")
	if data.FieldValue != "" {
		err = validation.Validate(r.Context(), data)
		if err != nil {
			data.Err = err.Error()
			data.Invalid = true
		} else {
			data.Valid = true
		}
	}

	utils.RenderTemplate(global.AppData, w, "validated-input-field", data)
}
