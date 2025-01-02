package config

import (
	"log"
	"net/http"
	"os"
)

type App struct {
	Server  *http.ServeMux
	InfoLog *log.Logger
	ErrLog  *log.Logger
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
