package api

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	val "github.com/Remi-Godin/simple-log/internal/utils/validation"
	"github.com/Remi-Godin/simple-log/internal/utils/validation/validators"
	"github.com/rs/zerolog/log"
)

type NewUserData struct {
	Email validators.EmailValidator
	Links map[string]string
}

func newNewUserData() NewUserData {
	return NewUserData{
		Email: validators.NewEmailValidator(),
		Links: make(map[string]string),
	}
}

func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	data := newNewUserData()
	data.Links["Submit"] = fmt.Sprintf("/register/user")
	data.Links["ValidatePassword"] = fmt.Sprintf("/register/validate/password")
	data.Links["ValidateEmail"] = fmt.Sprintf("/register/validate/email")
	data.Links["ValidateFirstName"] = fmt.Sprintf("/register/validate/first-name")
	data.Links["ValidateLastName"] = fmt.Sprintf("/register/validate/last-name")

	log.Info().Msg("IT WORKS HERE ANYWAY")

	data.Email.Links["ValidateEmail"] = fmt.Sprintf("/register/validate/email")
	data.Email.FieldValue = r.FormValue("email")
	err := val.Validate(r.Context(), data.Email)
	log.Info().Msg("IT WORKS THERE TOO")
	if err != nil {
		log.Err(err)
	}

	data.Email.Links["ValidateField"] = fmt.Sprintf("/register/validate/email")

	if r.Method == "GET" {
		utils.RenderTemplate(global.AppData, w, "register-form", data)
		return
	}

	fail := false
	err = r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	queryParams := database.InsertNewUserParams{}

	/*
			queryParams.Firstname = r.FormValue("first-name")
			data.FirstName.FieldData = r.FormValue("first-name")
			if len(queryParams.Firstname) < 2 {
				log.Error().Msg("First name too short")
				fail = true
				data.FirstName.Err = "First name must be longer than 1 character"
			}

			queryParams.Lastname = r.FormValue("last-name")
			data.LastName.FieldData = r.FormValue("last-name")
			if len(queryParams.Lastname) < 2 {
				log.Error().Msg("Last name too short")
				fail = true
				data.LastName.Err = "Last name must be longer than 1 character"
			}
		queryParams.Email = r.FormValue("email")
	*/

	/*
		password := r.FormValue("password")
		err = passwordValidator.Validate(password, 60)
		if err != nil {
			data.Password.Err = err.Error()
			fail = true
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fail = true
			data.Password.Err = err.Error()
		} else {
			queryParams.Passwordhash = string(hash)
		}
	*/

	if fail {
		w.WriteHeader(http.StatusBadRequest)
		utils.RenderTemplate(global.AppData, w, "register-form", data)
	} else {
		database.New(global.AppData.Conn).InsertNewUser(r.Context(), queryParams)
		utils.RenderTemplate(global.AppData, w, "success", data)
	}

}
