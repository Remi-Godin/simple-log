package api

import (
	"net/http"

	entries "github.com/Remi-Godin/simple-log/internal/api/entries"
	logbook "github.com/Remi-Godin/simple-log/internal/api/logbook"
	users "github.com/Remi-Godin/simple-log/internal/api/users"
	validation "github.com/Remi-Godin/simple-log/internal/api/validation"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func SetRoutes(mux *http.ServeMux) {
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("web/styles"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))

	mux.HandleFunc("/", Index)
	mux.HandleFunc("GET /logbook/{logbookId}/entries", entries.GetEntriesFromLogbook)
	mux.HandleFunc("GET /logbook/{logbookId}/entries/{entryId}", entries.GetEntryFromLogbook)
	mux.HandleFunc("POST /logbook/{logbookId}/entries", entries.InsertNewEntryInLogbook)
	mux.HandleFunc("DELETE /logbook/{logbookId}/entries/{entryId}", entries.DeleteEntryFromLogbook)
	mux.HandleFunc("PATCH /logbook/{logbookId}/entries/{entryId}", entries.UpdateEntryFromLogbook)

	mux.HandleFunc("GET /logbook/{logbookId}/modal/create", ModalCreate)
	mux.HandleFunc("GET /logbook/{logbookId}/entries/{entryId}/modal/edit", ModalEdit)

	mux.HandleFunc("DELETE /logbook/{logbookId}", logbook.DeleteLogbook)
	mux.HandleFunc("POST /logbook", logbook.InsertNewLogbook)
	mux.HandleFunc("GET /logbook", logbook.GetLogbooks)
	mux.HandleFunc("GET /logbook/{logbookId}", logbook.GetLogbook)

	mux.HandleFunc("GET /register/validate/password", validation.ValidatePasswordStrength)
	mux.HandleFunc("GET /register/validate/email", validation.ValidateEmail)
	mux.HandleFunc("GET /register/validate/first-name", validation.ValidateFirstName)
	mux.HandleFunc("GET /register/validate/last-name", validation.ValidateLastName)

	mux.HandleFunc("GET /register", Register)
	mux.HandleFunc("GET /register/user", users.InsertNewUser)
	mux.HandleFunc("POST /register/user", users.InsertNewUser)
}

type HomepageData struct {
	Links map[string]string
}

func newHomepageData() HomepageData {
	return HomepageData{
		Links: make(map[string]string),
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := newHomepageData()
	data.Links["Register"] = "/register"
	utils.RenderTemplate(global.AppData, w, "homepage", data)
}

func Register(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Redirecting to user registration page")
	http.Redirect(w, r, "/register/user", http.StatusSeeOther)
}
