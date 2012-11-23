// golanger project main.go
package main

import (
	"code.google.com/p/gorilla/mux"
	"db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	config map[string]string
	store  *sessions.CookieStore
)

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		println("read config file error")
		panic(err)
		os.Exit(1)
	}

	defer file.Close()

	dec := json.NewDecoder(file)

	err = dec.Decode(&config)

	if err != nil {
		println("read config file error")
		panic(err)
		os.Exit(1)
	}

	cookie_secret := config["cookie_secret"]

	store = sessions.NewCookieStore([]byte(cookie_secret))
}

func main() {

	http.Handle("/goweb/", http.FileServer(http.Dir("static")))

	router := mux.NewRouter()
	for url, handler := range handlers {
		router.HandleFunc(url, handler)
	}

	http.Handle("/", router)

	port := config["port"]

	println("Listen", port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

func indexPage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("IndexPage")
	t, _ := template.ParseFiles("./temp/Index.html")
	t.Execute(w, nil)
}

func registePage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("RegistePage")
	if req.Method == "GET" {
		t, _ := template.ParseFiles("./temp/Registe.html")
		t.Execute(w, nil)
	} else {
		req.ParseForm()
		username := req.FormValue("username")
		password := req.FormValue("password")
		cfpwd := req.FormValue("cfpwd")
		if password == cfpwd {
			db.InsertUser(username, password)
			http.Redirect(w, req, "/index", 301)
		}
		http.Redirect(w, req, "/registe", 301)
	}
}
