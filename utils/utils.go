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