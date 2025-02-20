package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"goth/internal/handlers"
)

/*
* Set to production at build time
* used to determine what assets to load
 */
var Environment = "development"

func init() {
	os.Setenv("env", Environment)
	// run generate script
	cmd := exec.Command("make", "tailwind-build")
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Register route handlers
	mux.HandleFunc("GET /{$}", handlers.NewHomeHandler().ServeHTTP)
	mux.HandleFunc("POST /{$}", handlers.NewPostHomeHandler().ServeHTTP)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Listening on port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		fmt.Println(err.Error())
	}
}
