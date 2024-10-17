package pages

import (
	"net/http"
	"time"

	"github.com/Remi-Godin/simple-log/internal/auth"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func LoginRedirect(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Redirecting to login page")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if auth.ValidateRequest(r, auth.NewSimpleJwtHandler(global.AppData.Env.AuthSecret, time.Minute)) {
		http.Redirect(w, r, "/logbook", http.StatusSeeOther)
		return
	}
	data := NewPageData("login")
	data.Links["LoginForm"] = "/form/login"
	data.Links["Register"] = "/page/register"
	utils.RenderTemplate(global.AppData, w, "page-login", data)
}
