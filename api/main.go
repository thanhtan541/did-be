package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/thanhtan541/did-be-wp/api/configuration"
	"github.com/thanhtan541/did-be-wp/api/startup"
)

func main() {
	config, err := configuration.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	application, err := startup.Build(config)
	if err != nil {
		log.Fatalf("❌ Failed to build app: %v", err)
	}

	log.Printf("Listening on port %d", application.Port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := application.RunUntilStopped(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
