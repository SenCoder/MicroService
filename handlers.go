package main

import (
	"net/http"
	"fmt"
	"time"

	"github.com/unrolled/render"
)

type sampleContent struct {
	ID string `json:"id"`
	Content string `json:"content"`
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
		sampleContent{ID:"8675309", Content:"Hello from Go!"})
	}
}


func helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Gorilla!\n"))
	}
}


func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data := sampleContent{ID:"8675309", Content:"Hello from Go!"}
		t.Execute(w, data)
	}
}


func cookieWriteHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		expiration := time.Now().Add(2 * 24 * time.Hour)
		cookie := http.Cookie{Name :"sample", Value: "this is a gorilla cookie", Expires: expiration}
		http.SetCookie(w, &cookie)
		formatter.JSON(w, http.StatusOK, "cookie set")
	}
}


func cookieReadHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		
		cookie, err := req.Cookie("sample")
		if err == nil {
			fmt.Fprint(w, cookie.Value)
		} else {
			fmt.Fprintf(w, "read cookie fail: %v", err)
		}
		
	}
}