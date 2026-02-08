package analyzer

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// link analysis results.
type LinkStats struct {
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
}

// AnalyzeLinks performs the link analysis for a given document and page URL.
// - Collect all links from the document
// - Extract links
// - Normalize links
// - Classify internal/external
// - Concurrent accessibility checks
// - Aggregate counts
// - Return stats
func AnalyzeLinks(ctx context.Context, pageURL string, doc *goquery.Document) (LinkStats, error) {
	var stats LinkStats

	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return stats, err
	}

	// collect all links from the document
	links := collectLinks(doc)

	// Concurrency Setup
	// wait until all link checks complete
	var wg sync.WaitGroup
	// protect shared counter, without mutex: race condition
	var mu sync.Mutex

	for _, link := range links {
		// convert relative to absolute URL
		// /about to https://example.com/about
		absoluteURL := resolveLink(baseURL, link)
		if absoluteURL == nil {
			continue
		}

		// classify internal vs external
		// same host -> internal
		// different host -> external
		if absoluteURL.Host == baseURL.Host {
			stats.InternalLinks++
		} else {
			stats.ExternalLinks++
		}

		// check accessibility concurrently
		// network calls are slow, so we check links in parallel to speed up the process.
		wg.Add(1)
		go func(link string) {
			defer wg.Done()

			if !isAccessible(ctx, link) {
				mu.Lock()
				stats.InaccessibleLinks++
				mu.Unlock()
			}
		}(absoluteURL.String())
	}

	wg.Wait()
	return stats, nil
}

// collectLinks collects href attributes from anchor tags.
func collectLinks(doc *goquery.Document) []string {
	var links []string

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		// get the href attribute's value
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		// remove leading and trailing white space
		href = strings.TrimSpace(href)
		if href == "" {
			return
		}

		// skip non http links
		if strings.HasPrefix(href, "mailto:") ||
			strings.HasPrefix(href, "tel:") ||
			strings.HasPrefix(href, "javascript:") {
			return
		}

		links = append(links, href)
	})

	return links
}

func resolveLink(base *url.URL, href string) *url.URL {
	u, err := url.Parse(href)
	if err != nil {
		return nil
	}

	return base.ResolveReference(u)
}

// extractLinks collects href attributes from anchor tags.
func extractLinks(doc *goquery.Document) []string {
	var links []string

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		// get the href attribute's value
		href, _ := s.Attr("href")

		if shouldIgnoreLink(href) {
			return
		}

		links = append(links, href)
	})

	return links
}

// shouldIgnoreLink filters out non-HTTP links.
func shouldIgnoreLink(href string) bool {
	href = strings.TrimSpace(href)
	return href == "" ||
		strings.HasPrefix(href, "mailto:") ||
		strings.HasPrefix(href, "tel:") ||
		strings.HasPrefix(href, "javascript:")
}

// isAccessible Checks whether a given link is reachable over HTTP
func isAccessible(ctx context.Context, link string) bool {
	// Timeout for network calls can hang
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	// Head method checks if resource exists and does not download the body
	request, error := http.NewRequestWithContext(ctx, http.MethodHead, link, nil)
	if error != nil {
		return false
	}

	// send request
	response, error := client.Do(request)
	if error != nil {
		return false
	}
	// close the responsebody
	defer response.Body.Close()

	// if status < 400, can Accessible
	return response.StatusCode < http.StatusBadRequest
}
