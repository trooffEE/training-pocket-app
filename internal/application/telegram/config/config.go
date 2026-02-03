package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	Token       string `env:"TOKEN"`
	Port        string `env:"PORT"`
	AdminId     string `env:"ADMIN_TELEGRAM_ID"`
	Development bool   `env:"DEVELOPMENT"`
}

func New() Config {
	if err := godotenv.Load(); err != nil {
		zap.L().Fatal("Error loading environment variables", zap.Error(err))
	}

	var config = Config{
		Token:       os.Getenv("TELEGRAM_TOKEN"),
		Port:        os.Getenv("TELEGRAM_WEB_SERVER_PORT"),
		AdminId:     os.Getenv("TELEGRAM_ADMIN_ID"),
		Development: os.Getenv("DEVELOPMENT") == "1",
	}

	return config
}
