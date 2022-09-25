package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"ubersnap/internal/config"
	"ubersnap/pkg/logger"

	"github.com/gorilla/mux"
)

var (
	cfg       = config.Server()
	appLogger = logger.NewApiLogger()
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)

	appLogger.Info("ubersnap Service Run on " + cfg.Port)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := http.ListenAndServe(cfg.Port, router)
		if err != nil {
			appLogger.Error(err)
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		appLogger.Error(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		appLogger.Error(fmt.Sprintf("ctx.Done: %v", done))
	}
}
