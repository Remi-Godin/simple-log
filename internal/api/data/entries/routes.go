package entries

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/auth"
)

func LoadEntriesRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.Handle("GET /logbook/{logbookId}/entries", auth.WithAuth(GetEntriesFromLogbook))
	mux.Handle("GET /logbook/{logbookId}/entries/{entryId}", auth.WithAuth(GetEntryFromLogbook))

	mux.Handle("POST /logbook/{logbookId}/entries", auth.WithAuth(InsertNewEntryInLogbook))
	mux.Handle("DELETE /logbook/{logbookId}/entries/{entryId}", auth.WithAuth(DeleteEntryFromLogbook))
	mux.Handle("PATCH /logbook/{logbookId}/entries/{entryId}", auth.WithAuth(UpdateEntryFromLogbook))
	return mux
}
