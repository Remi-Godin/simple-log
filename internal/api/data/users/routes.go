package users

import "net/http"

func LoadUsersRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("POST /register", InsertNewUser)
	return mux
}
