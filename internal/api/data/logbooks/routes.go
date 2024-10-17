package logbooks

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/auth"
)

func LoadLogbooksRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.Handle("DELETE /logbook/{logbookId}", auth.WithAuth(DeleteLogbook))
	mux.Handle("GET /logbooks", auth.WithAuth(GetLogbooksOwnedBy))
	return mux
}
