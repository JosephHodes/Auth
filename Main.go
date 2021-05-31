package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//note it only accepts encoded form data
func auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		auth := authentication{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 400)
		}
		err = json.Unmarshal(body, &auth)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 400)
		}
		if auth.Username == "" || auth.Password == "" {
			w.WriteHeader(401)
		} else {
			w.WriteHeader(200)

			log.Println(auth.Username)
			log.Println(auth.Password)
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
