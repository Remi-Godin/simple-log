package pages

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func RegisterRedirect(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Redirecting to user registration page")
	http.Redirect(w, r, "/register", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	data := newPageData("register")
	data.Links["RegisterForm"] = "/form/register"
	utils.RenderTemplate(global.AppData, w, "page-register", data)
}
