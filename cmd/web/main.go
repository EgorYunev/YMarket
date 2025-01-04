package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/EgorYunev/YMarket/config"
	_ "github.com/lib/pq"
)

type App struct {
	Server  *http.ServeMux
	InfoLog *log.Logger
	ErrLog  *log.Logger
	DB      *sql.DB
}

func main() {
	app := New()
	start(app)

	app.InfoLog.Printf("Connecting to data base %s", config.DBAdress)
	app.startDB(config.DBAdress)
	defer app.DB.Close()

	app.InfoLog.Printf("Starting http server on %s port", config.HTTPAdress)
	log.Fatal(http.ListenAndServe(config.HTTPAdress, app.Server))
}

func (a *App) startDB(dbArd string) {
	db, err := sql.Open("postgres", dbArd)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		a.ErrLog.Fatalf("Cannot ping data base: %s", dbArd)
	}
	a.DB = db
}

func New() *App {
	infLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := App{
		InfoLog: infLog,
		ErrLog:  errLog,
	}
	return &app
}
