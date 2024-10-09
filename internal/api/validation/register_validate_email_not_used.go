package api

import (
	"fmt"
	"net/http"
	"net/mail"

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

	data := NewFieldValidationData()
	data.Links["ValidateEmail"] = fmt.Sprintf("/register/validate/email")

	data.FieldData = r.FormValue("email")
	_, err = mail.ParseAddress(data.FieldData)
	if err != nil {
		data.Err = err.Error()
	}
	validator.FieldValue = data.FieldData
	err = validation.Validate(validator)

	data.Err = "Email address already in use. Please select a different one."
	utils.RenderTemplate(global.AppData, w, "register-email", data)

}
