package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

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
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webroot + "/assets")))
	mx.HandleFunc("/api/test", testHandler(formatter)).Methods("GET")
}