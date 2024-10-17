package pages

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func LogbookRedirect(w http.ResponseWriter, r *http.Request) {
	logbookIdStr := r.PathValue("logbookId")
	logbookUrl := fmt.Sprintf("/logbook/%s", logbookIdStr)
	log.Info().Msg(fmt.Sprintf("Redirecting to %s", logbookUrl))
	w.Header().Add("HX-Redirect", logbookUrl)
}

func Logbook(w http.ResponseWriter, r *http.Request) {
	logbookId := r.PathValue("logbookId")
	data := NewPageData("Logbook")
	data.Links["InitialLoad"] = fmt.Sprintf("/logbook/%s/entries?limit=5&offset=0", logbookId)
	data.Data["LogbookTitle"] = "This is my logbook"
	data.Data["LogbookDescription"] = "This is my logbook description"
	data.Links["Create"] = "/form/entry"
	utils.RenderTemplate(global.AppData, w, "page-logbook", data)
}
