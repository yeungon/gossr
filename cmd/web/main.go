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

func main() {
	env := config.NewEnv()
	pool, err := pgxpool.New(context.Background(), env.POSTGRES_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	config := config.NewApp(true, true)
	config.Queries = db.NewQueries(pool)
	config.Conn = pool

	logger := log.New(os.Stdout, "[api] ", log.LstdFlags)

	server := app.NewServer(config, logger)

	log.Printf("Server running at %s\n", env.APP_PORT)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
