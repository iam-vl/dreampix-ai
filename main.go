package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("HTTP_LISTEN_ADDR")
	router := chi.NewMux()
	slog.Info("Application running", "port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func initEverything() error {
	// err := godotenv.Load()
	// if err != nil {
	// 	return err
	// }
	// return nil
	return godotenv.Load()
}
