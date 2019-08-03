package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pyaesone17/golang-myanmar-tutorial/simple-rest-server-03/config"
)

// NewMainHandler is factory method for creating Handler
func NewMainHandler(config config.Config) *MainHandler {
	return &MainHandler{config}
}

// MainHandler is storing config dependency
type MainHandler struct {
	Config config.Config
}

// Invoke is responsible to handle the request of "/"
func (handler *MainHandler) Invoke(res http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"Hello":   "World",
		"Version": handler.Config.Version,
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
