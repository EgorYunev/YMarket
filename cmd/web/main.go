package main

import (
	"flag"
	"net/http"

	"github.com/EgorYunev/YMarket/config"
)

func main() {
	adr := flag.String("adr", ":8080", "Server http adress")
	app := config.New()
	start(app)
	app.InfoLog.Printf("Starting http server on %s port", *adr)

	http.ListenAndServe(*adr, app.Server)
}
