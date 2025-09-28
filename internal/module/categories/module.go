package categories

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yeungon/gossr/config"
	"github.com/yeungon/gossr/internal/module/categories/business"
	"github.com/yeungon/gossr/internal/module/categories/storage"
	"github.com/yeungon/gossr/internal/module/categories/transport"
)

func Init(logger *log.Logger, cfg *config.AppConfig) (http.Handler, *business.CategoryService) {
	repo := storage.NewCategoryPostgres(cfg.Queries.Categories)
	svc := business.NewCategoryService(repo)
	handler := transport.NewCategoryHandler(svc, cfg)

	r := chi.NewRouter()
	r.Get("/{id}", handler.GetCategory)
	r.Get("/", handler.ListCategories)

	return r, svc
}
