package handlers

import (
	"../data"
	"encoding/json"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	version := os.Getenv("VERSION")
	content := []string{"Api server testing, version: " + version}

	defaultResponse := data.DefaultResponse{Content: content}
	response, err := json.Marshal(defaultResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}
