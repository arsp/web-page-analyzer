package analyzer

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// FetchHTML retrieves the raw HTML content of the given URL.
func FetchHTML(ctx context.Context, targetURL string) (string, int, error) {
	// Create an http client with a fixed timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// create a new HTTP request
	request, error := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if error != nil {
		return "", 0, fmt.Errorf("failed to create request: %w", error)
	}

	// Execute the request.
	response, error := client.Do(request)
	if error != nil {
		return "", 0, fmt.Errorf("failed to fetch url: %w", error)
	}
	defer response.Body.Close()

	// Get the response body.
	body, error := io.ReadAll(response.Body)
	if error != nil {
		return "", response.StatusCode, fmt.Errorf("failed to read response body: %w", error)
	}

	// Handle non 200 responses as errors
	if response.StatusCode != http.StatusOK {
		return "", response.StatusCode, fmt.Errorf(
			"received non-OK HTTP status: %d %s",
			response.StatusCode,
			http.StatusText(response.StatusCode),
		)
	}

	return string(body), response.StatusCode, nil
}
