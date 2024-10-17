package pages

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
)

func Logbooks(w http.ResponseWriter, r *http.Request) {
	data := NewPageData("Logbook")
	data.Links["InitialLoad"] = fmt.Sprintf("/logbooks?limit=5&offset=0")
	utils.RenderTemplate(global.AppData, w, "page-logbook", data)
}
