package handlers

import "net/http"

// HealthHandler responds with a simple "OK" message to indicate the service is healthy
func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}

}
