package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/trooffEE/training-app/internal/application/telegram"
	"github.com/trooffEE/training-app/internal/application/telegram/config"
	"github.com/trooffEE/training-app/internal/application/telegram/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg := config.New()
	httpShutdown := server.Init(ctx, cfg)
	appShutdown := telegram.Start(ctx, cfg)

	<-ctx.Done()

	appShutdown()
	httpShutdown()
}
