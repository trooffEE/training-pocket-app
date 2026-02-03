package commander

import (
	"fmt"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"github.com/trooffEE/training-app/internal/application/telegram/commander/api"
	"github.com/trooffEE/training-app/internal/application/telegram/config"
	"go.uber.org/zap"
)

type Commander struct {
	API *api.Api
}

func New(cfg config.Config, bot *tgbotapi.BotAPI) *Commander {
	apiCmder, err := api.NewApi(cfg, bot)
	if err != nil {
		zap.S().Fatalw("Failed to create api interface", "error", err)
		return nil
	}

	return &Commander{
		API: apiCmder,
	}
}

func (c *Commander) Start(update tgbotapi.Update) {
	fmt.Println("Hello world!", update.Message.From.UserName)
}
