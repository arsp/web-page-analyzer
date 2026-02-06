package server

import (
	"html/template"
	"log/slog"
	"net/http"
	"web-page-analyzer/internal/validator"
)

// registerRoutes registers all HTTP endpoints
// exposed by the application.
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/analyze", analyzeHandler)
}

// indexHandler renders the home page that contains
// the URL input form.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and execute the HTML template.
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	tmpl.Execute(w, nil)
}

// analyzeHandler handles form submissions from the UI.
// extracts the URL provided by the user and triggers
func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	// Only POST requests are allowed for analysis.
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the URL from the submitted form.
	url := r.FormValue("url")

	// Validate the URL before processing
	if err := validator.ValidateURL(url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the incoming analysis request.
	slog.Info("received analyze request", "url", url)

	w.Write([]byte("Analysis started for: " + url))
}
