package main

import (
	"html/template"
	"net/http"

	"github.com/EgorYunev/YMarket/pkg/models"
)

func start(app *App) {
	serv := http.NewServeMux()

	serv.HandleFunc("/", app.homeHandler)
	serv.HandleFunc("/reg", app.registration)
	serv.HandleFunc("/user/add", app.addUser)
	serv.HandleFunc("/ad/search", app.getAdsFiltered)
	app.Server = serv
}

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {

	tm, err := template.ParseFiles("./ui/html/main.html")

	if err != nil {
		app.ErrLog.Printf("Cannot parse html files")
		http.Error(w, "Cannot parse html files", http.StatusInternalServerError)
		return
	}

	ads, err := app.Ads.GetLastest()
	if err != nil {
		app.ErrLog.Print(err)
	}

	tm.Execute(w, ads)
}

func account(w http.ResponseWriter, r *http.Request) {

}

func addAd(w http.ResponseWriter, r *http.Request) {

}

func deleteAd(w http.ResponseWriter, r *http.Request) {

}

func getAdById(w http.ResponseWriter, r *http.Request) {

}

func (app *App) getAdsFiltered(w http.ResponseWriter, r *http.Request) {

	tm, err := template.ParseFiles("./ui/html/ads.html")

	if err != nil {
		app.ErrLog.Print(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	name := r.FormValue("title")

	ads, err := app.Ads.GetAdsFiltered(name)

	if err != nil {
		app.ErrLog.Print(err)
		http.Error(w, "Cannot found any data", http.StatusInternalServerError)
		return
	}

	tm.Execute(w, ads)

}

func (app *App) addUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Wrond Method", http.StatusBadRequest)
		return
	}

	name := r.FormValue("username")
	pass := r.FormValue("password")

	err := app.Users.Insert(name, pass)

	if err != nil {
		app.ErrLog.Print(err)
		http.Error(w, "Cannot add new user", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func getUserById(w http.ResponseWriter, r *http.Request) {

}

func changeUser(w http.ResponseWriter, r *http.Request) {

}

func (app *App) registration(w http.ResponseWriter, r *http.Request) {

	tm, err := template.ParseFiles("./ui/html/reg.html")
	if err != nil {
		http.Error(w, "Ixternal server error", http.StatusInternalServerError)
		return
	}

	user := models.User{}

	tm.Execute(w, user)
}
