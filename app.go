// app.go

package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)
//App - References to the router and the database
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize - Initializes the application
func (a *App) Initialize(user, password, dbname string) { }

//Run - Runs the application
func (a *App) Run(addr string) { }