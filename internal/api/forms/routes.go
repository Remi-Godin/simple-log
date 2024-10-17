package forms

import "net/http"

func LoadFormsRoutes(mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc("GET /form/register", Register)
	mux.HandleFunc("GET /form/login", Login)
	mux.HandleFunc("GET /form/entry", NewEntry)
	mux.HandleFunc("GET /form/logbook/{logbookId}/entries/{entryId}", EditEntry)
	mux.HandleFunc("GET /form/logbook", Logbook)

	return mux
}
