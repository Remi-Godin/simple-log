package fields

import "net/http"

func LoadFieldsRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /field/validated-password", ValidatePasswordStrength)
	mux.HandleFunc("GET /field/validated-email", ValidateEmail)
	mux.HandleFunc("GET /field/validated-first-name", ValidateFirstName)
	mux.HandleFunc("GET /field/validated-last-name", ValidateLastName)
	mux.HandleFunc("GET /field/email", Email)
	mux.HandleFunc("GET /field/password", Password)
	return mux
}
