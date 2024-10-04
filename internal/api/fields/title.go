package fields

import (
	"net/http"
	"net/url"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func Title(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("There was an error when parsing the form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := NewGenericFormField()
	data.FieldName = "title"
	data.FieldId = "title"
	data.FieldNameText = "Title"
	data.FieldType = "text"
	data.FieldValue, _ = url.QueryUnescape(r.FormValue("value"))
	log.Warn().Msg(data.FieldValue)

	utils.RenderTemplate(global.AppData, w, "com-input-field", data)
}
