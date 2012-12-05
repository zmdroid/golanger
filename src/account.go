/*
*user login | register
 */
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//return current user from session
func currentUser(req *http.Request) (*User, bool) {

	session, _ := store.Get(req, "user")
	username, ok := session.Values["username"]

	if !ok {
		return nil, false
	}

	name, ok := username.(string)

	user := User{}
	if !user.findUser(name) {
		return nil, false
	}
	return &user, true
}

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
		user := User{}
		_, flag := user.vilUser(username, password)
		if flag {

			session, _ := store.Get(req, "user")
			session.Values["username"] = username
			session.Save(req, w)

			http.Redirect(w, req, "/admin", http.StatusFound)
		}

		http.Redirect(w, req, "/", http.StatusFound)

	}
}

//User register
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
			user := &User{}
			user.Username = username
			user.Userpwd = password
			if user.registeUser() {

				session, _ := store.Get(req, "golanger")
				session.Values["sessionid"] = username
				session.Save(req, w)

				http.Redirect(w, req, "/index", http.StatusFound)
			}
			http.Redirect(w, req, "/registe", http.StatusFound)
		}
		http.Redirect(w, req, "/registe", http.StatusFound)
	}
}
