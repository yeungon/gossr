package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yeungon/gossr/config"

	"github.com/yeungon/gossr/internal/module/articles"
)

func NewRouter(cfg *config.AppConfig, logger *log.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// init article module with its own sqlc
	articleHandler, _ := articles.Init(logger, cfg)

	r.Mount("/articles", articleHandler)

	return r
}
