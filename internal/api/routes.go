package api

import (
	"net/http"
	"time"

	entries "github.com/Remi-Godin/simple-log/internal/api/entries"
	forms "github.com/Remi-Godin/simple-log/internal/api/forms"
	fields "github.com/Remi-Godin/simple-log/internal/api/forms/fields"
	logbook "github.com/Remi-Godin/simple-log/internal/api/logbook"
	"github.com/Remi-Godin/simple-log/internal/api/pages"
	users "github.com/Remi-Godin/simple-log/internal/api/users"
	"github.com/Remi-Godin/simple-log/internal/auth"
	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func SetRoutes(mux *http.ServeMux) {
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("web/styles"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))

	// Page redirect handling
	mux.HandleFunc("/", pages.LoginRedirect)
	mux.HandleFunc("GET /page/register", pages.RegisterRedirect)

	mux.HandleFunc("GET /page/success", pages.SuccessRedirect)
	mux.HandleFunc("GET /success", pages.Success)

	// Form fields
	mux.HandleFunc("GET /field/validated-password", fields.ValidatePasswordStrength)
	mux.HandleFunc("GET /field/validated-email", fields.ValidateEmail)
	mux.HandleFunc("GET /field/validated-first-name", fields.ValidateFirstName)
	mux.HandleFunc("GET /field/validated-last-name", fields.ValidateLastName)
	mux.HandleFunc("GET /field/email", fields.LoginEmail)
	mux.HandleFunc("GET /field/password", fields.LoginPassword)

	mux.HandleFunc("GET /register", pages.Register)
	mux.HandleFunc("POST /register", users.InsertNewUser)

	mux.HandleFunc("GET /form/register", forms.Register)
	mux.HandleFunc("GET /form/login", forms.Login)

	mux.HandleFunc("POST /login", Login)
	mux.HandleFunc("GET /login", pages.Login)

	// SECURE ROUTES

	// Entries
	mux.Handle("GET /logbook/{logbookId}/entries", WithAuth(entries.GetEntriesFromLogbook))
	mux.Handle("GET /logbook/{logbookId}/entries/{entryId}", WithAuth(entries.GetEntryFromLogbook))
	mux.Handle("POST /logbook/{logbookId}/entries", WithAuth(entries.InsertNewEntryInLogbook))
	mux.Handle("DELETE /logbook/{logbookId}/entries/{entryId}", WithAuth(entries.DeleteEntryFromLogbook))
	mux.Handle("PATCH /logbook/{logbookId}/entries/{entryId}", WithAuth(entries.UpdateEntryFromLogbook))

	// Logbooks
	mux.Handle("DELETE /logbook/{logbookId}", WithAuth(logbook.DeleteLogbook))
	mux.Handle("POST /logbook", WithAuth(logbook.InsertNewLogbook))
	mux.Handle("GET /logbook", WithAuth(logbook.GetLogbooks))
	mux.Handle("GET /logbook/{logbookId}", WithAuth(logbook.GetLogbook))

	// Success
	mux.Handle("GET /secure/success", WithAuth(pages.Success))
}

// Processes login form and sends JWT on successful auth
func Login(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Processing login information...")

	// Parse form for data
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Could not parse login form.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Get password hash for associated email from database
	hash, err := database.New(global.AppData.Conn).GetUserPasswordHash(r.Context(), email)
	if err != nil {
		log.Error().Err(err).Msg("Could not find user.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Compare hash and password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Error().Err(err).Msg("Wrong password")
		w.WriteHeader(http.StatusForbidden)
		// TODO
		utils.RenderTemplate(global.AppData, w, "form-error", nil)
		return
	}

	// Create new authentication token
	jwtHandler := auth.NewSimpleJwtHandler(global.AppData.Env.AuthSecret, time.Minute)
	token, err := jwtHandler.GenerateToken(email)
	if err != nil {
		log.Error().Err(err).Msg("Could not generate token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Bind token to cookie and attach to response
	cookie := auth.BindJwtToCookie(token)
	http.SetCookie(w, cookie)
	//w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
	log.Info().Msg("I'm gonna redirect now, watch this!")

	// Redirect to success page (temporary)
	w.Header().Add("HX-Redirect", "/secure/success")
}

// Wrapper function to add Auth middleware to routes
func WithAuth(handleFunc http.HandlerFunc) http.Handler {
	return auth.AuthMiddleware(handleFunc, global.AppData.Env.AuthSecret)
}
