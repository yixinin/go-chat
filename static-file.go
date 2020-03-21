package main

import "net/http"

func init() {
	http.Handle("/static/avatar/", http.StripPrefix("/static/avatar/", http.FileServer(http.Dir("static/avatar"))))
}
