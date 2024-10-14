package api

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/database"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/Remi-Godin/simple-log/internal/utils/validation"
	"github.com/Remi-Godin/simple-log/internal/utils/validation/validators"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		links := make(map[string]string)
		links["ValidateFirstName"] = "/field/validated-first-name"
		links["ValidateLastName"] = "/field/validated-last-name"
		links["ValidateEmail"] = "/field/validated-email"
		links["ValidatePassword"] = "/field/validated-password"
		links["Submit"] = "/register"
		utils.RenderTemplate(global.AppData, w, "com-form", links)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var validationErrors []string
	firstName := validators.NewNameValidator()
	firstName.FieldValue = r.FormValue("first-name")
	lastName := validators.NewNameValidator()
	lastName.FieldValue = r.FormValue("last-name")
	email := validators.NewEmailValidator()
	email.FieldValue = r.FormValue("email")
	password := validators.NewPasswordValidator()
	password.FieldValue = r.FormValue("password")

	if err = validation.Validate(r.Context(), firstName); err != nil {
		validationErrors = append(validationErrors, fmt.Sprintf("First name: %s", err.Error()))
	}
	if err = validation.Validate(r.Context(), lastName); err != nil {
		validationErrors = append(validationErrors, fmt.Sprintf("Last name: %s", err.Error()))
	}
	if err = validation.Validate(r.Context(), email); err != nil {
		validationErrors = append(validationErrors, fmt.Sprintf("Email: %s", err.Error()))
	}
	if err = validation.Validate(r.Context(), password); err != nil {
		validationErrors = append(validationErrors, fmt.Sprintf("Password: %s", err.Error()))
	}

	if len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		utils.RenderTemplate(global.AppData, w, "form-submission-error", "Please fill all required fields and fix all issues")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password.FieldValue), bcrypt.DefaultCost)

	queryParams := database.InsertNewUserParams{
		Firstname:    firstName.FieldValue,
		Lastname:     lastName.FieldValue,
		Email:        email.FieldValue,
		Passwordhash: string(hash),
	}
	_, err = database.New(global.AppData.Conn).InsertNewUser(r.Context(), queryParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Add("HX-Redirect", "/success")
	w.WriteHeader(http.StatusCreated)

}
