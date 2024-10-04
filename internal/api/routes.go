package api

import (
	"net/http"
	"time"

	"github.com/Remi-Godin/simple-log/internal/api/data/entries"
	"github.com/Remi-Godin/simple-log/internal/api/data/logbooks"
	"github.com/Remi-Godin/simple-log/internal/api/data/users"
	"github.com/Remi-Godin/simple-log/internal/api/fields"
	"github.com/Remi-Godin/simple-log/internal/api/forms"
	"github.com/Remi-Godin/simple-log/internal/api/pages"
	"github.com/Remi-Godin/simple-log/internal/auth"
	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func SetRoutes(mux *http.ServeMux) {
	// Asset routes
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("web/styles"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))

	forms.LoadFormsRoutes(mux)
	pages.LoadPagesRoutes(mux)
	fields.LoadFieldsRoutes(mux)
	users.LoadUsersRoutes(mux)
	logbooks.LoadLogbooksRoutes(mux)
	entries.LoadEntriesRoutes(mux)

	mux.HandleFunc("POST /login", Login)
}

// Processes login form and sends JWT on successful auth
func Login(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Processing login information...")

	// Parse form for data
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Could not parse login form.")
		w.WriteHeader(http.StatusBadRequest)
		utils.RenderTemplate(global.AppData, w, "form-submission-error", "Could not process login information.")
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Get password hash for associated email from database
	hash, err := database.New(global.AppData.Conn).GetUserPasswordHash(r.Context(), email)
	if err != nil {
		log.Error().Err(err).Msg("Could not find user.")
		w.WriteHeader(http.StatusBadRequest)
		utils.RenderTemplate(global.AppData, w, "form-submission-error", "This email does not have an account associated with it.")
		return
	}

	// Compare hash and password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Error().Err(err).Msg("Wrong password")
		w.WriteHeader(http.StatusForbidden)
		// TODO
		utils.RenderTemplate(global.AppData, w, "form-submission-error", "Wrong password.")
		return
	}

	// Create new authentication token
	jwtHandler := auth.NewSimpleJwtHandler(global.AppData.Env.AuthSecret, time.Minute*60)
	token, err := jwtHandler.GenerateToken(email)
	if err != nil {
		log.Error().Err(err).Msg("Could not generate token")
		w.WriteHeader(http.StatusInternalServerError)
		utils.RenderTemplate(global.AppData, w, "form-submission-error", "The server encountered an error.")
		return
	}

	// Bind token to cookie and attach to response
	cookie := auth.BindJwtToCookie(token)
	http.SetCookie(w, cookie)
	//w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
	log.Info().Msg("I'm gonna redirect now, watch this!")

	// Redirect to success page (temporary)
	w.Header().Add("HX-Redirect", "/logbook")
}
