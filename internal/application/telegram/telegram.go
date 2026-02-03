package telegram

import (
	"context"
	"sync"
	"time"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
	"github.com/trooffEE/training-app/internal/application/telegram/commander"
	"github.com/trooffEE/training-app/internal/application/telegram/config"
	"go.uber.org/zap"
)

func Start(ctx context.Context, cfg config.Config) func() {
	botApi, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		zap.L().Error("Filed to create new bot api", zap.Error(err))
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := botApi.GetUpdatesChan(updateConfig)
	cmder := commander.New(cfg, botApi)

	var wg sync.WaitGroup
	wg.Go(func() { handleUpdate(updates, cmder) })

	return func() {
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		// graceful shutdown, check -> server.go
		wg.Wait()
	}
}

func handleUpdate(updates tgbotapi.UpdatesChannel, cmder *commander.Commander) {
	for update := range updates {
		if update.Message != nil {
			if !update.Message.From.IsBot {
				zap.L().Info(
					"client message",
					zap.String("msg", update.Message.Text),
					zap.String("username", update.Message.From.UserName),
				)
			}

			switch update.Message.Text {
			case "/start":
				cmder.Start(update)
			}
		}

		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			if _, err := cmder.API.Bot.Request(callback); err != nil {
				zap.L().Error("Error receiving response from callback with id", zap.Error(err), zap.String("id", update.CallbackQuery.ID))
				continue
			}

			switch update.CallbackQuery.Data {

			}
		}
	}
}
