// server.go
package app

import (
	"log"
	"net/http"

	"github.com/yeungon/gossr/config"
)

func NewServer(cfg *config.AppConfig, logger *log.Logger) *http.Server {
	r := NewRouter(cfg, logger)

	return &http.Server{
		Addr:    cfg.APP_DOMAIN_URL, // Use APP_DOMAIN_URL instead of APP_PORT
		Handler: r,
	}
}
