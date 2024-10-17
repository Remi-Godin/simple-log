package forms

import "net/http"

func LoadFormsRoutes(mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc("GET /form/register", Register)
	mux.HandleFunc("GET /form/login", Login)
	mux.HandleFunc("GET /form/entry", Entry)
	mux.HandleFunc("GET /form/logbook", Logbook)

	return mux
}
