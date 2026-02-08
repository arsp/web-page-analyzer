package analyzer

import (
	"context"
	"strings"
	"web-page-analyzer/internal/model"

	"github.com/PuerkitoBio/goquery"
)

// Analyze performs the complete analysis of a given URL.
// it will orchestrate fetching the HTML, parsing the content,
// checking links, and aggregating the results.
//
// context parameter allows cancellation and timeout control.
func Analyze(ctx context.Context, url string) (model.AnalysisResult, int, error) {
	// Step 1: Fetch HTML content
	html, status, err := FetchHTML(ctx, url)
	if err != nil {
		return model.AnalysisResult{}, status, err
	}

	// Step 2: Parse document structure
	title, version, headings, hasLogin, err := ParseHTML(html)
	if err != nil {
		return model.AnalysisResult{}, status, err
	}

	// Step 3: Create goquery document, needed for link analysis
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return model.AnalysisResult{}, status, err
	}

	// Step 4: Analyze links
	linkStats, err := AnalyzeLinks(ctx, url, doc)
	if err != nil {
		return model.AnalysisResult{}, status, err
	}

	// Step 5: Aggregate result
	result := model.AnalysisResult{
		URL:               url,
		HTMLVersion:       version,
		Title:             title,
		Headings:          headings,
		InternalLinks:     linkStats.InternalLinks,
		ExternalLinks:     linkStats.ExternalLinks,
		InaccessibleLinks: linkStats.InaccessibleLinks,
		HasLoginForm:      hasLogin,
	}

	return result, status, nil
}
