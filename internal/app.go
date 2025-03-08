package internal

import (
	"context"
	"fmt"
	"github.com/alishojaeiir/GoShort/config"
	"github.com/alishojaeiir/GoShort/internal/http"
	httpserver "github.com/alishojaeiir/GoShort/pkg/http_server"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Application struct {
	Config     config.Config
	HTTPServer http.Server
}

func Setup(ctx context.Context, cfg config.Config) Application {
	httpServer := http.New(*httpserver.New(cfg.Server))
	return Application{
		Config:     cfg,
		HTTPServer: httpServer,
	}
}

func (app Application) Start() {
	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	startServers(app, &wg)
	<-ctx.Done()
	fmt.Println("Shutdown signal received...")

	shutdownTimeoutCtx, cancel := context.WithTimeout(context.Background(), app.Config.TotalShutdownTimeout)
	defer cancel()

	if app.shutdownServers(shutdownTimeoutCtx) {
		fmt.Println("Servers shut down gracefully")
	} else {
		fmt.Println("Shutdown timed out, exiting application")
		os.Exit(1)
	}
	fmt.Println("app stopped")
}

func startServers(app Application, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(fmt.Sprintf("HTTP server started on %d", app.Config.Server.Port))
		if err := app.HTTPServer.Serve(); err != nil {
			fmt.Sprintln("error in HTTP server on %d", app.Config.Server.Port)
		}
		fmt.Sprintln("HTTP server stopped %d", app.Config.Server.Port)
	}()

}

func (app Application) shutdownServers(ctx context.Context) bool {
	shutdownDone := make(chan struct{})

	go func() {
		var shutdownWg sync.WaitGroup
		shutdownWg.Add(1)
		go app.shutdownHTTPServer(&shutdownWg)

		shutdownWg.Wait()
		close(shutdownDone)
	}()

	select {
	case <-shutdownDone:
		return true
	case <-ctx.Done():
		return false
	}
}

func (app Application) shutdownHTTPServer(wg *sync.WaitGroup) {
	defer wg.Done()
	httpShutdownCtx, httpCancel := context.WithTimeout(context.Background(), app.Config.Server.ShutDownCtxTimeout)
	defer httpCancel()
	if err := app.HTTPServer.Stop(httpShutdownCtx); err != nil {
		fmt.Sprintln("HTTP server graceful shutdown failed: %v", err)
	}
}
