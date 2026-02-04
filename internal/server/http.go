package server

import (
	"log/slog"
	"net/http"
	"os"
)

// Start initializes and starts the HTTP server.
// This function is responsible only for server lifecycle
func Start() {
	// Initialize a structured logger using slog
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Create a new HTTP request multiplexer.
	// This will be used to register all application routes.
	mux := http.NewServeMux()
	registerRoutes(mux)

	slog.Info("starting server", "port", 8080)

	// Start the HTTP server.
	// ListenAndServe blocks the current goroutine.
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("server failed to start", "error", err)
	}
}
