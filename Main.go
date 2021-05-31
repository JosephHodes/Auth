package main

import (
	"fmt"
	"net/http"
)

//note it only accepts encoded form data
func auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(200)
			fmt.Println(username)
			fmt.Println(password)
		}

	}
}

func handleRequests() {
	http.HandleFunc("/", auth)
	http.ListenAndServe(":8000", nil)
}
func main() {

	handleRequests()
}
