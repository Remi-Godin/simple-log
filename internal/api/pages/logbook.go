package pages

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Remi-Godin/simple-log/internal/database"
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
	logbookIdStr := r.PathValue("logbookId")
	logbookId, err := strconv.Atoi(logbookIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logbookData, err := database.New(global.AppData.Conn).GetLogbookData(r.Context(), int32(logbookId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := NewPageData("Logbook")
	data.Links["InitialLoad"] = fmt.Sprintf("/logbook/%d/entries?limit=5&offset=0", logbookId)
	data.Data["LogbookTitle"] = logbookData.Title
	data.Data["LogbookDescription"] = logbookData.Description
	data.Links["Create"] = fmt.Sprintf("/form/logbook/%d/entries", logbookId)
	data.Links["DeleteLogbook"] = fmt.Sprintf("/logbook/%d", logbookId)
	utils.RenderTemplate(global.AppData, w, "page-logbook", data)
}
