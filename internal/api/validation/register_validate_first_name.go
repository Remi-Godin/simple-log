package api

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func ValidateFirstName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := NewFieldValidationData()
	data.Links["ValidateFirstName"] = fmt.Sprintf("/register/validate/first-name")

	data.FieldData = r.FormValue("first-name")
	if len(data.FieldData) < 2 {
		data.Err = "First name must be longer than 1 character."
	}

	utils.RenderTemplate(global.AppData, w, "register-first-name", data)

}
