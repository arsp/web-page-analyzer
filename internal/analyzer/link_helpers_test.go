package analyzer

import (
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestResolveLink_InvalidURL(t *testing.T) {
	base, _ := url.Parse("https://example.com")
	// Malformed URL that url.Parse will reject
	u := resolveLink(base, "http://%41:8080/") // invalid host format

	if u != nil {
		t.Errorf("expected nil for invalid URL, got %v", u)
	}
}

func TestCollectLinks_IgnoresNonHTTP(t *testing.T) {
	html := `
	<html><body>
		<a href="mailto:test@example.com">Mail</a>
		<a href="javascript:void(0)">JS</a>
		<a href="tel:+123456">Tel</a>
		<a href="">Empty</a>
		<a href="/ok">OK</a>
	</body></html>
	`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	links := collectLinks(doc)

	if len(links) != 1 || links[0] != "/ok" {
		t.Errorf("expected only '/ok', got %#v", links)
	}
}
