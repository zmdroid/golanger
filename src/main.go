// golanger project main.go
package main

import (
	"code.google.com/p/gorilla/mux"
	"encoding/json"
	"github.com/gorilla/sessions"
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
	http.Handle("/scripts/", http.FileServer(http.Dir("static")))
	http.Handle("/content/", http.FileServer(http.Dir("static")))
	http.Handle("/images/", http.FileServer(http.Dir("static")))
	http.Handle("/Index.html", http.FileServer(http.Dir("temp")))

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
