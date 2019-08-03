package handler

import (
	"encoding/json"
	"net/http"
)

// MainHandler is responsible to handle the request of "/"
func MainHandler(res http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"Hello": "World",
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
