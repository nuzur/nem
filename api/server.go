package main

import (
	"net/http"

	"go.uber.org/config"
	"go.uber.org/zap"
)

func httpServer(config config.Provider, logger *zap.Logger) {
	// http port from config
	httpPort := config.Get("ports.http").String()

	// http upload - move to seperate module
	http.Handle("/upload", serverHandlerUploadFunc())

	go http.ListenAndServe(":"+httpPort, nil)

	logger.Info(`Serving HTTP on PORT: %s`, zap.String("port", httpPort))
}
