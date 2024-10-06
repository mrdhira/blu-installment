package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	installmentCtrl "blu-installment/controller/installment"
	"blu-installment/router"
	installmentService "blu-installment/service/installment"
)

var (
	author string
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	// Define command-line flags
	flag.StringVar(&author, "author", "Dhira Wigata", "author name for copyright attribution")

	// Parse the command-line flags
	flag.Parse()

	installmentSvc := installmentService.New(logger)
	installmentCtrl := installmentCtrl.New(logger, installmentSvc)

	mux := router.New(installmentCtrl)

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	// Start the HTTP server
	go func() {
		logger.Info(fmt.Sprintf("Starting server on port %s", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed to start", slog.String("error", err.Error()))
		}
	}()

	// Graceful shutdown
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// Use a buffered channel to avoid missing signals as the graceful
	// shutdown process can take a long time.
	quitSignal := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-quitSignal

	// Create a deadline to wait for.
	// Wait for 5 seconds before shutting down the server.
	logger.Info("Shutting down server...")
	if err := server.Shutdown(context.TODO()); err != nil {
		logger.Error("failed to shutdown http server", slog.String("error", err.Error()))
		os.Exit(1)
	}
	logger.Info("Server gracefully stopped")

	os.Exit(0)

}
