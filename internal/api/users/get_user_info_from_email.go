package api

import (
	"net/http"
)

type EmailInputData struct {
	Email       string
	AlreadyUsed bool
}

func GetUserInfoFromEmail(w http.ResponseWriter, r *http.Request) {

}
