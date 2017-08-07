package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var t *template.Template


func init() {
	t = template.Must(template.ParseFiles("assets/templates/index.html"))
}

func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	mx := mux.NewRouter()
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webroot := os.Getenv("WEBROOT")
	if len(webroot) == 0 {
		root, err := os.Getwd()
		if err != nil {
			panic("Could not retrive working directory")
		} else {
			webroot = root
		}
	}

	mx.HandleFunc("/", homeHandler()).Methods("GET")
	mx.HandleFunc("/cookies/write", cookieWriteHandler(formatter)).Methods("GET")
	mx.HandleFunc("/cookies/read", cookieReadHandler(formatter)).Methods("GET")
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webroot + "/assets")))
	
}



