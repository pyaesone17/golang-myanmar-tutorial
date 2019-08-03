package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pyaesone17/golang-myanmar-tutorial/simple-rest-server-04/config"
)

// MainHandler is responsible to handle the request of "/"
func MainHandler(res http.ResponseWriter, req *http.Request, conf config.Config) {
	data := map[string]interface{}{
		"Hello":   "World",
		"Version": conf.Version,
	}

	bytes, err := json.Marshal(data)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Something went wrong"))
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(bytes)
}
