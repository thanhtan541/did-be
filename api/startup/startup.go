package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhtan541/did-be-wp/api/configuration"
	"github.com/thanhtan541/did-be-wp/api/route"
)

type Application struct {
	Port   int
	Server *http.Server
}

type ApplicationBaseUrl string

func Build(cfg *configuration.Settings) (*Application, error) {
	address := fmt.Sprintf("%s:%d", cfg.Application.Host, cfg.Application.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to bind to address %s: %w", address, err)
	}

	port := listener.Addr().(*net.TCPAddr).Port

	router := gin.New()
	router.GET("/ping", route.Ping)
	srv := &http.Server{
		Handler: router,
	}

	go func() {
		log.Printf("🚀 Starting server on %s", listener.Addr())
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	return &Application{
		Port:   port,
		Server: srv,
	}, nil
}

func (a *Application) RunUntilStopped(ctx context.Context) error {
	<-ctx.Done()
	log.Println("🔌 Shutting down server...")
	return a.Server.Shutdown(context.Background())
}
