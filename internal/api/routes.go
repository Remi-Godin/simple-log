package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func SetRoutes(mux *http.ServeMux) {
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("web/styles"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))

	mux.HandleFunc("/", Index)
	mux.HandleFunc("GET /logbook", GetLogbooks)
	mux.HandleFunc("GET /logbook/{logbookId}", GetLogbook)
	mux.HandleFunc("GET /logbook/{logbookId}/entries", GetEntriesFromLogbook)
	mux.HandleFunc("GET /logbook/{logbookId}/entries/{entryId}", GetEntryFromLogbook)
	mux.HandleFunc("POST /logbook/{logbookId}/entries", InsertNewEntryInLogbook)
	mux.HandleFunc("POST /logbook", InsertNewLogbook)
	mux.HandleFunc("DELETE /logbook/{logbookId}/entries/{entryId}", DeleteEntryFromLogbook)
	mux.HandleFunc("DELETE /logbook/{logbookId}", DeleteLogbook)
	mux.HandleFunc("GET /modal/create", ModalCreate)
	mux.HandleFunc("GET /modal/edit/{logbookId}/{entryId}", ModalEdit)
	mux.HandleFunc("PATCH /logbook/{logbookId}/entries/{entryId}", UpdateEntryFromLogbook)

}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Yup, this is the index")
}
