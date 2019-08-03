package main

import (
	"net/http"

	"github.com/pyaesone17/golang-myanmar-tutorial/simple-rest-server-03/config"
	"github.com/pyaesone17/golang-myanmar-tutorial/simple-rest-server-03/handler"
)

func main() {
	appconfig := config.NewConfig()
	mainHandler := handler.NewMainHandler(appconfig)

	http.HandleFunc("/", mainHandler.Invoke)
	http.ListenAndServe(":8080", nil)
}
