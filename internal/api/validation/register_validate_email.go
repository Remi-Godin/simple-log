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

	validator := validators.NewEmailValidator()
	validator.FieldValue = r.FormValue("email")
	validator.Links["ValidateField"] = fmt.Sprintf("/register/validate/email")
	validator.FieldValue = r.FormValue("email")
	err = validation.Validate(r.Context(), validator)
	if err != nil {
		log.Err(err)
		validator.Err = err.Error()
	}

	utils.RenderTemplate(global.AppData, w, "validated-input-field", validator)

}
