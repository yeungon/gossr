package articles

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yeungon/gossr/config"
	"github.com/yeungon/gossr/internal/module/articles/business"
	"github.com/yeungon/gossr/internal/module/articles/storage"
	transport "github.com/yeungon/gossr/internal/module/articles/transport/http"
)

func Init(logger *log.Logger, cfg *config.AppConfig) (http.Handler, *business.ArticleService) {
	repo := storage.NewArticlePostgres(cfg.Queries.Articles)
	svc := business.NewArticleService(repo)
	handler := transport.NewArticleHandler(svc, cfg)

	r := chi.NewRouter()
	r.Get("/article/{id}", handler.GetArticle) //

	r.Get("/view/{test}", handler.GetMostViewArticle) //
	return r, svc
}
