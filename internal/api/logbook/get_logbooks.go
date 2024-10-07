package api

import (
	"encoding/json"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/rs/zerolog/log"
)

func GetLogbooks(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := database.New(global.AppData.Conn).GetLogbooksOwnedBy(r.Context(), r.FormValue("email"))
	if err != nil {
		log.Error().Err(err).Msg("Could not complete database query")
		w.WriteHeader(http.StatusInternalServerError)
	}

	enc := json.NewEncoder(w)
	enc.Encode(result)
}
