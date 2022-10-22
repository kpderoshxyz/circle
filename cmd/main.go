package main

import (
	"fmt"
	"net/http"

	"github.com/kpderoshxyz/circle/api"
)

func main() {
	fmt.Printf("Starting server")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	handler, err := api.NewHandler()
	if err != nil {
		panic(err)
	}
	mux.Handle("/clock", handler)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error closing the server")
	}
	fmt.Printf("Application is done")
}
