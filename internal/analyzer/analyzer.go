package analyzer

import "context"

// Analyze performs the complete analysis of a given URL.
// it will orchestrate fetching the HTML, parsing the content,
// checking links, and aggregating the results.
//
// context parameter allows cancellation and timeout control.
func Analyze(ctx context.Context, url string) error {
	// 1. Fetch HTML content
	html, _, err := FetchHTML(ctx, url)
	if err != nil {
		return err
	}

	// 2. Parse document structure
	_, _, _, _, err = ParseHTML(html)
	if err != nil {
		return err
	}

	// 3. Analyze links
	// 4. Detect login forms
	// 5. Return aggregated result
	return nil
}
