package middleware

import "net/http"

func GetKey(r *http.Request) string {

	// Check API key first
	apiKey := r.Header.Get("X-API-Key")

	if apiKey != "" {
		return apiKey
	}

	// Otherwise use IP address
	return r.RemoteAddr
}
