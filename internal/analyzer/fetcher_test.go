package analyzer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestFetchHTMLSuccess verifies that HTML is fetched correctly
// when the server responds with HTTP 200.
func TestFetchHTMLSuccess(t *testing.T) {
	// Create a fake HTTP server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><title>Test</title></html>"))
	}))
	defer server.Close()

	ctx := context.Background()
	body, status, error := FetchHTML(ctx, server.URL)

	if error != nil {
		t.Fatalf("expected no error, got %v", error)
	}

	if status != http.StatusOK {
		t.Fatalf("expected status 200, got %d", status)
	}

	if body == "" {
		t.Fatal("expected non-empty body")
	}
}

// TestFetchHTMLNonOK verifies behavior for non-200 responses.
func TestFetchHTMLNonOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	ctx := context.Background()
	_, status, error := FetchHTML(ctx, server.URL)

	if error == nil {
		t.Fatal("expected error for non-OK status")
	}

	if status != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", status)
	}
}

// TestFetchHTMLTimeout verifies that context cancellation is respected.
func TestFetchHTMLTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	}))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, _, err := FetchHTML(ctx, server.URL)

	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
}
