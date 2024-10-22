package handler

import (
	"io"
	"net/http"
)

func PostInventory(w http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)
	if err != nil {
		ErrorResponseJSON(w, http.StatusInternalServerError, "Unable to read the provided data.")
		return
	}
}

func GetAllInventory(w http.ResponseWriter, r *http.Request) {
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
}

func PutInventory(w http.ResponseWriter, r *http.Request) {
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
}
