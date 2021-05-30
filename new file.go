package main

import (
	"fmt"
	"net/http"
)

func auth() {

}
func handlerequests(w *http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", auth)
}
func main() {

	handlerequests()
}
