package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trooffEE/training-app/internal/application/telegram/server"
	"github.com/trooffEE/training-app/internal/lib"
)

func main() {
	server := server.NewServer()

	done := make(chan bool, 1)

	go lib.GracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}
