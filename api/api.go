package api

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool `json:"success"`
	Body interface{} `json:"body"`
}

func (r *ApiResponse) RespondJSON(w http.ResponseWriter, status int) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(r)

	if err != nil {
		return
	}
}
