package main

import "net/http"

var (
	handlers = map[string]func(http.ResponseWriter, *http.Request){
		"/":           indexPage,
		"/login":      loginPage,
		"/index":      indexPage,
		"/registe":    registePage,
		"/admin":      adminManagerPage,
		"/newBlog":    newBlogPage,
		"/deleteBlog": deleteBlog,
		"/editBlog":   editBlog,
		"/page":       blogPage,
	}
)
