package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	adr := flag.String("adr", ":8080", "Server http adress")
	dbAdr := flag.String("dbadr", "root:admin@tcp(localhost:33060)/ymarket", "Database adress")
	app := New()
	start(app)

	app.InfoLog.Printf("Connecting to data base %s", *dbAdr)
	app.DB = startDB(*dbAdr)
	defer app.DB.Close()

	app.InfoLog.Printf("Starting http server on %s port", *adr)
	log.Fatal(http.ListenAndServe(*adr, app.Server))
}

func startDB(dbard string) *sql.DB {
	db, err := sql.Open("mysql", dbard)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

type App struct {
	Server  *http.ServeMux
	InfoLog *log.Logger
	ErrLog  *log.Logger
	DB      *sql.DB
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
