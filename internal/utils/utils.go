package utils

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	_ "github.com/Remi-Godin/simple-log/internal/database"
	"github.com/rs/zerolog/log"
)

type Env struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	DbAddr           string
	DbPort           string
	Port             string
	AuthSecret       string
}

type AppData struct {
	Conn *sql.DB
	Tmpl *template.Template
	Env  Env
}

type Link struct {
	Path   string
	Rel    string
	Method string
}

func (link Link) ToString() string {
	return fmt.Sprintf("%s %s", link.Method, link.Path)
}

func NewLink(path string, rel string, method string) Link {
	return Link{
		path,
		rel,
		method,
	}
}

func GenerateAppData(conn *sql.DB, tmpl *template.Template) AppData {
	data := AppData{
		Conn: conn,
		Tmpl: tmpl,
		Env:  *LoadEnvVars(),
	}
	return data
}

func LoadEnvVars() *Env {
	env := Env{}

	env.PostgresUser = os.Getenv("POSTGRES_USER")
	env.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	env.PostgresDb = os.Getenv("POSTGRES_DB")
	env.DbAddr = os.Getenv("DB_ADDR")
	env.DbPort = os.Getenv("DB_PORT")
	env.Port = os.Getenv("PORT")
	env.AuthSecret = os.Getenv("SECRET")

	return &env
}

func RenderTemplate(appData AppData, w http.ResponseWriter, tmpl_name string, data any) {
	log.Info().Msg(fmt.Sprint("Rendering template: ", tmpl_name))
	err := appData.Tmpl.ExecuteTemplate(w, tmpl_name, data)
	if err != nil {
		log.Error().Err(err).Msg("Could not execute template")
	}
}

func ExtractIdsFromRoute(r *http.Request) (logbookId int, entryId int, err error) {
	logbookId, err = strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Could not extract logbook ID from path")
		return 0, 0, err
	}
	entryId, err = strconv.Atoi(r.PathValue("entryId"))
	if err != nil {
		log.Error().Err(err).Msg("Could not extract entry ID from path")
		return 0, 0, err
	}
	return logbookId, entryId, nil
}
