package config

import (
	"log"
	"os"
	"text/template"

	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nalgeon/redka"
	"github.com/yeungon/gossr/internal/infra/db"
)

type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	ErrorLog       *log.Logger
	InfoLog        *log.Logger
	InProduction   bool
	AUTH_USER      string
	AUTH_PASSWORD  string
	APP_DOMAIN_URL string
	SessionManager *scs.SessionManager
	CSRFKey        []byte
	CSRFSecure     bool

	Queries *db.Queries   // ✅ wrapped sqlc queries
	Conn    *pgxpool.Pool // ✅ db connection pool
	Redka   *redka.DB
}

func NewApp(cacheState bool, ProductionState bool) *AppConfig {
	return &AppConfig{
		UseCache:     cacheState,
		InProduction: ProductionState,
	}

}

// type Config struct {
// 	HTTPAddr string
// 	DBURL    string
// }

// func Load() *Config {
// 	return &Config{
// 		HTTPAddr: envOr("HTTP_ADDR", ":8080"),
// 		DBURL:    envOr("DB_URL", "postgres://user:pass@localhost:5432/mydb?sslmode=disable"),
// 	}
// }

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
