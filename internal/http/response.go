package http

import (
	"encoding/json"
	"net/http"
)

type GenericResponse struct {
	Success bool `json:"success"`
	Body interface{} `json:"body"`
	Message string `json:"message"`
}

func (r *GenericResponse) RespondJSON(w http.ResponseWriter, status int) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(r)

	if err != nil {
		return
	}
}
