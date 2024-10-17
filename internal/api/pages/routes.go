package pages

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/auth"
)

func LoadPagesRoutes(mux *http.ServeMux) *http.ServeMux {
	// Public
	mux.HandleFunc("/", LoginRedirect)
	mux.HandleFunc("GET /page/register", RegisterRedirect)
	mux.HandleFunc("GET /page/success", SuccessRedirect)

	mux.HandleFunc("GET /success", Success)
	mux.HandleFunc("GET /register", Register)
	mux.HandleFunc("GET /login", Login)
	// Secure
	mux.Handle("GET /page/logbook/{logbookId}", auth.WithAuth(LogbookRedirect))

	mux.Handle("GET /logbook", auth.WithAuth(Logbooks))
	mux.Handle("GET /logbook/{logbookId}", auth.WithAuth(Logbook))
	mux.Handle("GET /secure/success", auth.WithAuth(Success))
	return mux
}
