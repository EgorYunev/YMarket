package main

import (
	"html/template"
	"net/http"
)

func start(app *App) {
	serv := http.NewServeMux()

	serv.HandleFunc("/", app.homeHandler)
	app.Server = serv
}

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {

	tm, err := template.ParseFiles("./ui/html/main.html")

	if err != nil {
		app.ErrLog.Printf("Cannot parse html files")
		http.Error(w, "Cannot parse html files", http.StatusInternalServerError)
	}

	tm.Execute(w, nil)
}

func account(w http.ResponseWriter, r *http.Request) {

}

func addAd(w http.ResponseWriter, r *http.Request) {

}

func deleteAd(w http.ResponseWriter, r *http.Request) {

}

func getAdById(w http.ResponseWriter, r *http.Request) {

}

func getAdsFiltered(w http.ResponseWriter, r *http.Request) {

}

func addUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func getUserById(w http.ResponseWriter, r *http.Request) {

}

func changeUser(w http.ResponseWriter, r *http.Request) {

}
