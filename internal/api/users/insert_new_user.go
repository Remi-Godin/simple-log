package api

import (
	"fmt"
	"net/http"
	"net/mail"

	validation "github.com/Remi-Godin/simple-log/internal/api/validation"
	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type NewUserData struct {
	FirstName validation.FieldValidationData
	LastName  validation.FieldValidationData
	Email     validation.FieldValidationData
	Password  validation.FieldValidationData
	Links     map[string]string
}

func newNewUserData() NewUserData {
	return NewUserData{
		FirstName: validation.NewFieldValidationData(),
		LastName:  validation.NewFieldValidationData(),
		Email:     validation.NewFieldValidationData(),
		Password:  validation.NewFieldValidationData(),
		Links:     make(map[string]string),
	}
}

func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	data := newNewUserData()
	data.Links["Submit"] = fmt.Sprintf("/register/user")
	data.Links["ValidatePassword"] = fmt.Sprintf("/register/validate/password")
	data.Links["ValidateEmail"] = fmt.Sprintf("/register/validate/email")
	data.Links["ValidateFirstName"] = fmt.Sprintf("/register/validate/first-name")
	data.Links["ValidateLastName"] = fmt.Sprintf("/register/validate/last-name")

	data.Password.Links["ValidatePassword"] = fmt.Sprintf("/register/validate/password")
	data.Email.Links["ValidateEmail"] = fmt.Sprintf("/register/validate/email")
	data.FirstName.Links["ValidateFirstName"] = fmt.Sprintf("/register/validate/first-name")
	data.LastName.Links["ValidateLastName"] = fmt.Sprintf("/register/validate/last-name")

	if r.Method == "GET" {
		utils.RenderTemplate(global.AppData, w, "register-form", data)
		return
	}

	fail := false
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	queryParams := database.InsertNewUserParams{}

	queryParams.Firstname = r.FormValue("first-name")
	if len(queryParams.Firstname) <= 1 {
		fail = true
		data.FirstName.Err = "First name must be longer than 1 character"
	}

	queryParams.Lastname = r.FormValue("last-name")
	if len(queryParams.Lastname) <= 1 {
		fail = true
		data.LastName.Err = "Last name must be longer than 1 character"
	}

	queryParams.Email = r.FormValue("email")
	_, err = mail.ParseAddress(queryParams.Email)
	if err != nil {
		fail = true
		data.Email.Err = err.Error()
	}

	password := r.FormValue("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fail = true
		data.Password.Err = err.Error()
	} else {
		queryParams.Passwordhash = string(hash)
	}

	log.Info().Msg(fmt.Sprintf("%s %s %s %s", queryParams.Firstname, queryParams.Lastname, queryParams.Email, hash))
	if fail {
		log.Info().Msg("ERROR FOUND")
		utils.RenderTemplate(global.AppData, w, "register-form", data)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		database.New(global.AppData.Conn).InsertNewUser(r.Context(), queryParams)
	}

}
