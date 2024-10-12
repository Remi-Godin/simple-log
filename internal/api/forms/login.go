package forms

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	data := NewFormData("Login")
	data.FormDesc = "Enter your SimpleLog account details."

	data.FormSubmissionLink = "/login"
	data.FormFields = append(data.FormFields, "/field/email")
	data.FormFields = append(data.FormFields, "/field/password")

	utils.RenderTemplate(global.AppData, w, "com-form", data)
}
