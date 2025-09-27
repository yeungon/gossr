package config

import "os"

type Config struct {
	HTTPAddr string
	DBURL    string
}

func Load() *Config {
	return &Config{
		HTTPAddr: envOr("HTTP_ADDR", ":8080"),
		DBURL:    envOr("DB_URL", "postgres://user:pass@localhost:5432/mydb?sslmode=disable"),
	}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
