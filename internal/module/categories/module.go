package order

import (
	"log"
	"net/http"

	"github.com/yeungon/gossr/internal/module/order/business"
	"github.com/yeungon/gossr/internal/module/order/storage"
	"github.com/yeungon/gossr/internal/module/order/transport"

	"github.com/yeungon/gossr/config"
	"github.com/yeungon/gossr/internal/infra/db/sqlc"
	itembiz "github.com/yeungon/gossr/internal/module/item/business"
)

func Init(logger *log.Logger, cfg *config.Config, itemSvc *itembiz.Service, q *sqlc.Queries) (http.Handler, *business.Service) {
	repo := storage.NewPostgresRepo(q)
	svc := business.NewService(repo, itemSvc)
	handler := transport.NewHandler(svc)
	return handler, svc
}
