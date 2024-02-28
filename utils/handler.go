package utils

import (
	"log"
	"net/http"
)

// Handler
func Handler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Access logs
		log.Printf("%s %s %s\n", r.Method, r.URL.Path, r.UserAgent())
		f(w, r)
	}
}
