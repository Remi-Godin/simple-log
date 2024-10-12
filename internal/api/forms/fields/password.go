package fields

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func LoginPassword(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("There was an error when parsing the form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := NewGenericFormField()
	data.FieldName = "password"
	data.FieldId = "password"
	data.FieldNameText = "Password"
	data.FieldType = "password"

	utils.RenderTemplate(global.AppData, w, "com-input-field", data)
}
