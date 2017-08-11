package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	"github.com/unrolled/render"
)

type sampleContent struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			sampleContent{ID: "8675309", Content: "Hello from Go!"})
	}
}

//func (a *App) initializeRoutes() {
//    a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
//    a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
//    a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
//    a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
//    a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
//}

//func helloHandler() http.HandlerFunc {
//	return func(w http.ResponseWriter, req *http.Request) {
//		w.Write([]byte("Gorilla!\n"))
//	}
//}

type User struct {
	Username string
	Password string
}

func getLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("==getLogin==")
		data := sampleContent{ID: "8675309", Content: "Hello from Go!"}
		t := template.Must(template.ParseFiles("assets/templates/login.html"))
		t.Execute(w, data)
	}
}

func postLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		u := User{Username: req.FormValue("username"), Password: req.FormValue("password")}

		if "" == u.Username {
			respondWithJSON(w, http.StatusOK, map[string]interface{}{"code": 0, "message": "请填写用户名"})
		}

		if "" == u.Password {
			respondWithJSON(w, http.StatusOK, map[string]interface{}{"code": 0, "message": "请填写密码"})
		}
		//		err, users := LoginUser(username, password)
		respondWithJSON(w, http.StatusOK, map[string]interface{}{"code": 1, "message": "贺喜你，登录成功"})

		fmt.Println("user:", u)
	}

}

func getMain() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("==getMain==")
		data := sampleContent{ID: "8675309", Content: "Hello from Go!"}
		t := template.Must(template.ParseFiles("assets/templates/login.html"))
		t.Execute(w, data)
	}
}

func cookieWriteHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		expiration := time.Now().Add(2 * 24 * time.Hour)
		cookie := http.Cookie{Name: "sample", Value: "this is a gorilla cookie", Expires: expiration}
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

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Body(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body.Close()
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
