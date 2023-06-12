package middleware

import (
	"log"
	"net/http"
)

func middleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request -> [%s] %q", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
