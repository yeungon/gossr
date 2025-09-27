package articles

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yeungon/gossr/config"
	"github.com/yeungon/gossr/internal/infra/db/sqlc"
	"github.com/yeungon/gossr/internal/module/articles/business"
	"github.com/yeungon/gossr/internal/module/articles/storage"
	"github.com/yeungon/gossr/internal/module/articles/transport"
)

func Init(logger *log.Logger, cfg *config.Config, q *sqlc.Queries) (http.Handler, *business.ArticleService) {
	repo := storage.NewArticlePostgres(q)
	svc := business.NewArticleService(repo, logger)
	handler := transport.NewArticleHandler(svc)

	r := chi.NewRouter()
	r.Get("/{id}", handler.GetArticle) // e.g. GET /items/123?id=123

	return r, svc
}
