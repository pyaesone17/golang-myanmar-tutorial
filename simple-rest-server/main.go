package main

import (
	"encoding/json"
	"net/http"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"Hello": "World",
	}
	if bytes, err := json.Marshal(data); err == nil {
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(bytes)
		return
	}
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte("Something went wrong"))
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}
