package app

import (
	"log"
	"net/http"

	"github.com/yeungon/gossr/config"
)

func NewServer(cfg *config.Config, logger *log.Logger) *http.Server {
	mux := NewRouter(cfg, logger)

	return &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: mux,
	}
}
