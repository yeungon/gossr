package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yeungon/gossr/config"
	"github.com/yeungon/gossr/internal/app"
)

func main() {
	cfg := config.Load()
	logger := log.New(os.Stdout, "[api] ", log.LstdFlags)

	server := app.NewServer(cfg, logger)

	log.Printf("Server running at %s\n", cfg.HTTPAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
