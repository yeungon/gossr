package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yeungon/gossr/config"

	"github.com/yeungon/gossr/internal/infra/db/sqlc"
	"github.com/yeungon/gossr/internal/module/articles"
	order "github.com/yeungon/gossr/internal/module/categories"
)

func NewRouter(cfg *config.Config, logger *log.Logger) http.Handler {
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		logger.Fatal(err)
	}

	q := sqlc.New(db)

	itemHandler, itemSvc := articles.Init(logger, cfg, q)
	orderHandler, _ := order.Init(logger, cfg, itemSvc, q)

	r := chi.NewRouter()

	//common middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Mount handlers
	r.Mount("/items", itemHandler)
	r.Mount("/orders", orderHandler)

	return r
}
