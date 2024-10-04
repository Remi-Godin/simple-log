package pages

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func SuccessRedirect(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Redirecting to success page")
	w.Header().Add("HX-Redirect", "/success")
}

func Success(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(global.AppData, w, "page-success", nil)
}
