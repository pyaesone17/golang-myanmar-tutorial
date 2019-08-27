package main

import (
	"net/http"

	"github.com/pyaesone17/golang-myanmar-tutorial/simple-rest-server-05/handler"
)

func main() {
	http.HandleFunc("/", handler.MainHandler)
	http.HandleFunc("/ok", handler.OkHandler)

	http.ListenAndServe(":8080", nil)
}
