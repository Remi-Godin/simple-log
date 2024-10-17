package forms

import (
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
)

func Logbook(w http.ResponseWriter, r *http.Request) {
	data := NewFormData("Create Logbook")
	data.FormDesc = "Enter logbook details."

	data.FormSubmissionLink = "/logbook"
	data.FormFields = append(data.FormFields, "/field/title")
	data.FormFields = append(data.FormFields, "/field/description")

	utils.RenderTemplate(global.AppData, w, "com-form", data)
}
