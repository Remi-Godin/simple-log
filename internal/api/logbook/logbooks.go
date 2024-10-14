package api

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/api/pages"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
)

func Logbooks(w http.ResponseWriter, r *http.Request) {
	data := pages.NewPageData("Logbook")
	user := "regodin@proton.me"
	data.Links["InitialLoad"] = fmt.Sprintf("/data/logbook/%s?limit=5&offset=0", user)
	utils.RenderTemplate(global.AppData, w, "page-logbook", data)
}
