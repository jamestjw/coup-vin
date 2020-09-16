package main

import (
	"encoding/json"
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, payload interface{}) {
	v, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(v))
}