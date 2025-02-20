package main

import (
	"fmt"
	"net/http"
	"os"

	"goth/internal/handlers"
)

/*
* Set to production at build time
* used to determine what assets to load
 */
var Environment = "development"

func init() {
	os.Setenv("env", Environment)
}

func disableCacheInDevMode(next http.Handler) http.Handler {
	if Environment != "development" {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("/static/",
		disableCacheInDevMode(
			http.StripPrefix("/static/",
				http.FileServer(http.Dir("./static")))))

	// Register route handlers
	mux.HandleFunc("GET /{$}", handlers.NewHomeHandler().ServeHTTP)
	mux.HandleFunc("POST /{$}", handlers.NewPostHomeHandler().ServeHTTP)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Println("Listening on port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		fmt.Println(err.Error())
	}
}
