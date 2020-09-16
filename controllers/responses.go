package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func errorResponse(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		jsonResponse(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	badRequestResponse(w)
}

func errorResponseWithMessage(w http.ResponseWriter, statusCode int, errMessage string) {
	jsonResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: errMessage,
	})
	return
}

func badRequestResponse(w http.ResponseWriter) {
	jsonResponse(w, http.StatusBadRequest, nil)
}

func internalServerErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func recordNotFoundResponse(w http.ResponseWriter) {
	errorResponseWithMessage(w, http.StatusNotFound, "Room not found.")
}
