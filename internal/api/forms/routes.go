package forms

import "net/http"

func LoadFormsRoutes(mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc("GET /form/register", Register)
	mux.HandleFunc("GET /form/login", Login)

	return mux
}
