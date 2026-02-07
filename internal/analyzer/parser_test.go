package analyzer

import "testing"

func TestParseHTML(t *testing.T) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Test Page</title>
	</head>
	<body>
		<h1>Main Heading</h1>
		<h2>Sub Heading 1</h2>
		<h2>Sub Heading 2</h2>

		<form>
			<input type="password" name="password">
		</form>
	</body>
	</html>
	`

	title, version, headings, hasLogin, err := ParseHTML(html)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if title != "Test Page" {
		t.Errorf("expected title 'Test Page', got '%s'", title)
	}

	if version != "HTML5" {
		t.Errorf("expected HTML5, got '%s'", version)
	}

	if headings["h1"] != 1 {
		t.Errorf("expected 1 h1, got %d", headings["h1"])
	}

	if headings["h2"] != 2 {
		t.Errorf("expected 2 h2, got %d", headings["h2"])
	}

	if !hasLogin {
		t.Error("expected login form to be detected")
	}
}
