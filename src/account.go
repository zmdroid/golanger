/*
*user login | register
 */
package main

import (
	"db"
	"fmt"
	"html/template"
	"net/http"
)

//User Login
func loginPage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("LoginPage")
	if req.Method == "GET" {
		t, _ := template.ParseFiles("./temp/Login.html")
		t.Execute(w, nil)
	} else {
		req.ParseForm()
		username := req.FormValue("username")
		password := req.FormValue("password")
		pwd := db.QueryUser(username)
		if pwd == password {
			session, _ := store.Get(req, "user")
			session.Values["username"] = username
			session.Save(req, w)

			http.Redirect(w, req, "/index", http.StatusFound)
		}

		http.Redirect(w, req, "/", http.StatusFound)

	}
}
