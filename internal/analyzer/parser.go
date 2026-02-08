package analyzer

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParseHTML analyzes raw HTML content and extracts
// structural information required by the application.
// First Parentheses ( html string ) -> input Parameters
// Second Parentheses ( title string, htmlVersion string, headings map[string]int, hasLoginForm bool, err error ) -> Return Values
func ParseHTML(html string) (title string, htmlVersion string, headings map[string]int, hasLoginForm bool, error error) {
	// Initialize heading counters
	headings = make(map[string]int)

	// Detect HTML version from doctype
	htmlVersion = detectHTMLVersion(html)

	// Parse HTML using goquery
	doc, error := goquery.NewDocumentFromReader(strings.NewReader(html))
	if error != nil {
		return "", "", nil, false, error
	}

	// Extract page title
	title = strings.TrimSpace(doc.Find("title").First().Text())

	// Count heading tags (h1–h6)
	for i := 1; i <= 6; i++ {
		tag := "h" + string(rune('0'+i))
		headings[tag] = doc.Find(tag).Length()
	}

	// find the login form
	// assumed if an input[type="password"] exists inside a form.
	hasLoginForm = doc.Find("form input[type='password']").Length() > 0

	return title, htmlVersion, headings, hasLoginForm, nil
}

// detectHTMLVersion inspects the DOCTYPE declaration
func detectHTMLVersion(html string) string {
	upperHtml := strings.ToUpper(html)

	switch {
	case strings.Contains(upperHtml, "<!DOCTYPE HTML>"):
		return "HTML5"
	case strings.Contains(upperHtml, "XHTML"):
		return "XHTML"
	case strings.Contains(upperHtml, "HTML 4.01"):
		return "HTML 4.01"
	case strings.Contains(upperHtml, "<!DOCTYPE"):
		return "Unknown HTML (Doctype Present)"
	default:
		return "Unknown (No Doctype)"
	}
}
