package models

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	SQLDB *sql.DB
	Env   string
}

type Application struct {
	Router *mux.Router
}
