package api

import (
	"fmt"
	"net/http"

	"github.com/Remi-Godin/simple-log/internal/api/pages"
	"github.com/Remi-Godin/simple-log/internal/global"
	"github.com/Remi-Godin/simple-log/internal/utils"
	"github.com/rs/zerolog/log"
)

func LogbookRedirect(w http.ResponseWriter, r *http.Request) {
	logbookIdStr := r.PathValue("logbookId")
	logbookUrl := fmt.Sprintf("/logbook/%s", logbookIdStr)
	log.Info().Msg(fmt.Sprintf("Redirecting to %s", logbookUrl))
	//http.Redirect(w, r, logbookUrl, http.StatusSeeOther)
	w.Header().Add("HX-Redirect", logbookUrl)
}

func GetLogbook(w http.ResponseWriter, r *http.Request) {
	logbookId := r.PathValue("logbookId")
	data := pages.NewPageData("Logbook")
	//user := "regodin@proton.me"
	data.Links["InitialLoad"] = fmt.Sprintf("/data/logbook/%s/entries?limit=5&offset=0", logbookId)
	utils.RenderTemplate(global.AppData, w, "page-logbook", data)
}
