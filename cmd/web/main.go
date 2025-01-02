package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/EgorYunev/YMarket/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	adr := flag.String("adr", ":8080", "Server http adress")
	dbAdr := flag.String("dbadr", "root:admin@tcp(localhost:33060)/ymarket", "Database adress")
	app := config.New()
	start(app)
	app.InfoLog.Printf("Connecting to data base %s", *dbAdr)
	startDB(*dbAdr)
	app.InfoLog.Printf("Starting http server on %s port", *adr)

	log.Fatal(http.ListenAndServe(*adr, app.Server))
}

func startDB(dbard string) *sql.DB {
	db, err := sql.Open("mysql", dbard)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
