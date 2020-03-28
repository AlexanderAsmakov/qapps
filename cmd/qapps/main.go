package main

import (
	"github.com/AlexanderAsmakov/qapps/internal/handlers"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/apps", handlers.GetAllApps)

	http.ListenAndServe(":"+port, mux)
}
