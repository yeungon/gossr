package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yeungon/gossr/config"
	"github.com/yeungon/gossr/internal/app"
	"github.com/yeungon/gossr/internal/infra/db"
)

// main.go
func main() {
	env := config.NewEnv()
	pool, err := pgxpool.New(context.Background(), env.POSTGRES_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	appConfig := config.NewApp(true, true)
	appConfig.Queries = db.NewQueries(pool)
	appConfig.Conn = pool
	appConfig.APP_PORT = env.APP_PORT
	appConfig.APP_DOMAIN_URL = env.APP_DOMAIN_URL // used for logging / external URLs

	logger := log.New(os.Stdout, "[api] ", log.LstdFlags)

	server := app.NewServer(appConfig, logger)

	log.Printf("Server running at %s\n", appConfig.APP_DOMAIN_URL)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
