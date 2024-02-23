package utils

import (
	"log"
	"net/http"
)

// Access logs
func Logger(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s\n", r.Method, r.URL.Path, r.UserAgent())
        f(w, r)
    }
}