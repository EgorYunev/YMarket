package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/EgorYunev/YMarket/config"
	"github.com/EgorYunev/YMarket/pkg/database"
	_ "github.com/lib/pq"
)

type App struct {
	Server  *http.ServeMux
	InfoLog *log.Logger
	ErrLog  *log.Logger
	Users   *database.UserModel
	Ads     *database.AdModel
}

func main() {
	app := New()
	start(app)

	app.InfoLog.Printf("Connecting to data base %s", config.DBAdress)
	db := app.startDB(config.DBAdress)

	defer db.Close()

	app.Users.DB = db
	app.Ads.DB = db

	app.InfoLog.Printf("Starting http server on %s port", config.HTTPAdress)
	log.Fatal(http.ListenAndServe(config.HTTPAdress, app.Server))
}

func (a *App) startDB(dbArd string) *sql.DB {
	db, err := sql.Open("postgres", dbArd)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		a.ErrLog.Fatalf("Cannot ping data base: %s", dbArd)
	}
	return db
}

func New() *App {
	infLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := App{
		InfoLog: infLog,
		ErrLog:  errLog,
		Users:   &database.UserModel{},
		Ads:     &database.AdModel{},
	}
	return &app
}
