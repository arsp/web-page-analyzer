package analyzer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestAnalyzeLinks(t *testing.T) {
	internalServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer internalServer.Close()

	externalServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer externalServer.Close()

	html := `
	<html>
	<body>
		<a href="/internal">Internal</a>
		<a href="` + externalServer.URL + `">External</a>
	</body>
	</html>
	`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	stats, err := AnalyzeLinks(context.Background(), internalServer.URL, doc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stats.InternalLinks != 1 {
		t.Errorf("expected 1 internal link, got %d", stats.InternalLinks)
	}

	if stats.ExternalLinks != 1 {
		t.Errorf("expected 1 external link, got %d", stats.ExternalLinks)
	}

	if stats.InaccessibleLinks != 1 {
		t.Errorf("expected 1 inaccessible link, got %d", stats.InaccessibleLinks)
	}
}
