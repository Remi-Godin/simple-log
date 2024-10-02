package utils

import (
	"database/sql"
	"html/template"
	"net/http"
	"os"
	"strconv"

	_ "github.com/Remi-Godin/simple-log/internal/database"
	"github.com/rs/zerolog/log"
)

type Env struct {
	Postgres_user     string
	Postgres_password string
	Postgres_db       string
	Db_addr           string
	Db_port           string
	Port              string
}

type AppData struct {
	Conn *sql.DB
	Tmpl *template.Template
	Env  Env
}

func GenerateAppData(conn *sql.DB, tmpl *template.Template) AppData {
	return AppData{
		Conn: conn,
		Tmpl: tmpl,
		Env:  *LoadEnvVars(),
	}
}

func LoadEnvVars() *Env {
	env := Env{}

	env.Postgres_user = os.Getenv("POSTGRES_USER")
	env.Postgres_password = os.Getenv("POSTGRES_PASSWORD")
	env.Postgres_db = os.Getenv("POSTGRES_DB")
	env.Db_addr = os.Getenv("DB_ADDR")
	env.Db_port = os.Getenv("DB_PORT")
	env.Port = os.Getenv("PORT")

	return &env
}

func RenderTemplate(appData AppData, w http.ResponseWriter, tmpl_name string, data any) {
	log.Info().Msg("Rendering template")
	err := appData.Tmpl.ExecuteTemplate(w, tmpl_name, data)
	if err != nil {
		log.Error().Err(err).Msg("Could not execute template")
	}
}

func ExtractIdsFromRoute(r *http.Request) (logbookId int, entryId int, err error) {
	logbookId, err = strconv.Atoi(r.PathValue("logbookId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		return 0, 0, err
	}
	entryId, err = strconv.Atoi(r.PathValue("entryId"))
	if err != nil {
		log.Error().Err(err).Msg("Attempted to use API with erroneous parameters")
		return 0, 0, err
	}
	return logbookId, entryId, nil
}
