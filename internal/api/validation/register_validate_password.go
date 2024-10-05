package api

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
	passwordValidator "github.com/wagslane/go-password-validator"
)

func ValidatePasswordStrength(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := NewFieldValidationData()
	data.Links["ValidatePassword"] = fmt.Sprintf("/register/validate/password")

	password := r.FormValue("password")
	data.FieldData = password
	err = passwordValidator.Validate(password, 60)
	if err != nil {
		data.Err = err.Error()
	}
	utils.RenderTemplate(global.AppData, w, "register-password", data)
}
