package utils

import (
	"net/http"
	"strings"
)

func GetAuthHeaderToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	headerParts := strings.Fields(authHeader)

	if len(headerParts) < 2 {
		return ""
	}

	return headerParts[1]
}

func HandlePreFlight(w http.ResponseWriter, r *http.Request) bool {
	if(r.Method != "OPTIONS") {
		return false
	}

	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.WriteHeader(http.StatusOK)
	return true
}