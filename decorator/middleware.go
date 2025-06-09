package decorator

import (
	"fmt"
	"log"
	"net/http"
)

// handler
func FetchString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// middleware : logging
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming Request:", r.Method, r.URL.Path)
		next(w, r)
	}
}

// middleware: auth
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer rahasia123" {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func InitMiddleware() {
	http.HandleFunc("/hellow", LoggingMiddleware(AuthMiddleware(FetchString)))

	log.Println("Server running at http://localhost:3030")
	http.ListenAndServe(":3030", nil)
}
