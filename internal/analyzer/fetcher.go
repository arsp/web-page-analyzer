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
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create request: %w", err)
	}

	// Execute the request.
	response, err := client.Do(request)
	if err != nil {
		// return "", 0, fmt.Errorf("failed to fetch url: %w", err)
		// This is where unreachable host/DNS/network errors occur
		return "", 0, fmt.Errorf("unable to reach host: %w", err)
	}
	defer response.Body.Close()

	// Get the response body.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", response.StatusCode, fmt.Errorf("failed to read response body: %w", err)
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
