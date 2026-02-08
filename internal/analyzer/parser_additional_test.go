package analyzer

import "testing"

func TestParseHTML_NoLoginForm(t *testing.T) {

	html := `
	<html>
	<head><title>No Login</title></head>
	<body>
		<h1>Header</h1>
	</body>
	</html>
	`

	title, version, headings, hasLogin, err := ParseHTML(html)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if title != "No Login" {
		t.Errorf("unexpected title")
	}

	if version == "" {
		t.Errorf("expected detected html version")
	}

	if hasLogin {
		t.Errorf("should not detect login form")
	}

	if headings["h1"] != 1 {
		t.Errorf("expected 1 h1")
	}
}
