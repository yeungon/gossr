package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once //
	env  *ENV      //
)

type ENV struct {
	APP_NAME                        string
	APP_DOMAIN_URL                  string
	APP_PORT                        string
	POSTGRES_URL                    string
	CLOUDFLARE_CAPTCHA              string
	SETTING_MULTIPLE_LOGIN          bool
	IN_PRODUCTION                   bool
	COOKIE_SECURE                   bool
	SALT                            string
	CDN_URL                         string
	TELEGRAM_SEALQUESTIONS_CHANNEL  string
	TELEGRAM_SEALQUESTIONS_BOTTOKEN string
}

// Register the config
func NewEnv() *ENV {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		production := os.Getenv("IN_PRODUCTION")
		COOKIE_SECURE := os.Getenv("COOKIE_SECURE")
		POSTGRES_URL := os.Getenv("POSTGRES_URL")
		APP_DOMAIN_URL := os.Getenv("APP_DOMAIN_URL")
		APP_NAME := os.Getenv("APP_NAME")
		APP_PORT := os.Getenv("APP_PORT")
		CLOUDFLARE_CAPTCHA := os.Getenv("CLOUDFLARE_CAPTCHA")
		SETTING_MULTIPLE_LOGIN := os.Getenv("SETTING_MULTIPLE_LOGIN")
		CDN_URL := os.Getenv("CDN_URL")
		TELEGRAM_SEALQUESTIONS_CHANNEL := os.Getenv("TELEGRAM_SEALQUESTIONS_CHANNEL")
		TELEGRAM_SEALQUESTIONS_BOTTOKEN := os.Getenv("TELEGRAM_SEALQUESTIONS_BOTTOKEN")

		salt := os.Getenv("SALT")
		env = &ENV{
			SALT:                            salt,
			IN_PRODUCTION:                   convertStringToBool(production),
			COOKIE_SECURE:                   convertStringToBool(COOKIE_SECURE),
			POSTGRES_URL:                    POSTGRES_URL,
			APP_DOMAIN_URL:                  APP_DOMAIN_URL,
			APP_PORT:                        APP_PORT,
			APP_NAME:                        APP_NAME,
			CLOUDFLARE_CAPTCHA:              CLOUDFLARE_CAPTCHA,
			SETTING_MULTIPLE_LOGIN:          convertStringToBool(SETTING_MULTIPLE_LOGIN),
			CDN_URL:                         CDN_URL,
			TELEGRAM_SEALQUESTIONS_CHANNEL:  TELEGRAM_SEALQUESTIONS_CHANNEL,
			TELEGRAM_SEALQUESTIONS_BOTTOKEN: TELEGRAM_SEALQUESTIONS_BOTTOKEN,
		}
	})
	return env
}

func convertStringToBool(env_state string) bool {
	if len(env_state) < 1 {
		return false
	}
	bool, err := strconv.ParseBool(env_state)
	if err != nil {
		return false
	}
	return bool
}

func Get() *ENV {
	return env
}
