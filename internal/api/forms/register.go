package forms

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	data := NewFormData("Register")
	data.FormDesc = "Enter your details to create a SimpleLog account."

	data.FormSubmissionLink = "/register"
	data.FormFields = append(data.FormFields, "/field/validated-first-name")
	data.FormFields = append(data.FormFields, "/field/validated-last-name")
	data.FormFields = append(data.FormFields, "/field/validated-email")
	data.FormFields = append(data.FormFields, "/field/validated-password")

	utils.RenderTemplate(global.AppData, w, "com-form", data)
}
