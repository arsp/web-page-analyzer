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
func Analyze(ctx context.Context, url string) (model.AnalysisResult, error) {
	// Step 1: Fetch HTML content
	html, _, error := FetchHTML(ctx, url)
	if error != nil {
		return model.AnalysisResult{}, error
	}

	// Step 2: Parse document structure
	title, version, headings, hasLogin, error := ParseHTML(html)
	if error != nil {
		return model.AnalysisResult{}, error
	}

	// Step 3: Create goquery document, needed for link analysis
	doc, error := goquery.NewDocumentFromReader(strings.NewReader(html))
	if error != nil {
		return model.AnalysisResult{}, error
	}

	// Step 4: Analyze links
	linkStats, error := AnalyzeLinks(ctx, url, doc)
	if error != nil {
		return model.AnalysisResult{}, error
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

	return result, nil
}
