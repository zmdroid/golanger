package main

import "net/http"

var (
	handlers = map[string]func(http.ResponseWriter, *http.Request){
		"/":        loginPage,
		"/login":   loginPage,
		"/index":   indexPage,
		"/registe": registePage,
	}
)
