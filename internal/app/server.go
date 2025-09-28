package app

import (
	"log"
	"net/http"

	"github.com/yeungon/gossr/config"
)

func NewServer(cfg *config.AppConfig, logger *log.Logger) *http.Server {
	mux := NewRouter(cfg, logger)

	return &http.Server{
		Addr:    cfg.APP_DOMAIN_URL,
		Handler: mux,
	}
}
