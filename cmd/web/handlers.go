package main

import (
	"net/http"

	"github.com/EgorYunev/YMarket/config"
)

func start(app *config.App) {
	serv := http.NewServeMux()

	serv.HandleFunc("/", homeHandler)
	app.Server = serv
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
