package main

import (
	"fmt"
	"net/http"
)

func handlerequests(w *http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/")
}
func main() {

	handlerequests()
}
