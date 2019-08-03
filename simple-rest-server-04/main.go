package main

import (
	"net/http"

	"github.com/pyaesone17/golang-myanmar-tutorial/simple-rest-server-04/config"
	"github.com/pyaesone17/golang-myanmar-tutorial/simple-rest-server-04/handler"
)

// APIHandlerFunc is type for general handler
type APIHandlerFunc func(res http.ResponseWriter, req *http.Request, conf config.Config)

func dependencyInjector(handler APIHandlerFunc, appconfig config.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, appconfig)
	})
}

func main() {
	appconfig := config.NewConfig()
	http.HandleFunc("/", dependencyInjector(handler.MainHandler, appconfig))
	http.ListenAndServe(":8080", nil)
}
