package main

import (
	"log"
	"net/http"
	"os"
)

type Settings = SettingsSchemaJson
type Expenses = ExpensesSchemaJson

func main() {
	addr := env("ADDR", ":8080")

	srv, err := NewServer(ServerConfig{
		WebFS:    webFS,
		ConfigFS: configFS,
		DataFS:   dataFS,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server started and listening on %s", addr)

	if err := http.ListenAndServe(addr, srv.Routes()); err != nil {
		log.Fatal(err)
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
