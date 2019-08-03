package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pyaesone17/golang-myanmar-tutorial/simple-rest-server-05/config"
)

// MainHandler is responsible to handle the request of "/"
func MainHandler(res http.ResponseWriter, req *http.Request) {
	conf := config.NewConfig()
	data := map[string]interface{}{
		"Hello":   "World",
		"Version": conf.Version,
	}

	fmt.Printf("%p\n", conf)

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
