package analyzer

import "testing"

func TestParseHTML_HTML401_NoHeadings_NoLogin(t *testing.T) {
	html := `
	<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN">
	<html>
	<head>
		<title>Legacy Page</title>
	</head>
	<body>
		<p>No headings here</p>
	</body>
	</html>
	`

	title, version, headings, hasLogin, err := ParseHTML(html)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if title != "Legacy Page" {
		t.Errorf("expected title 'Legacy Page', got %q", title)
	}

	if version != "HTML 4.01" {
		t.Errorf("expected version HTML 4.01, got %q", version)
	}

	// Ensure heading counts default to 0 when absent
	for i := 1; i <= 6; i++ {
		tag := "h" + string(rune('0'+i))
		if headings[tag] != 0 {
			t.Errorf("expected %s count 0, got %d", tag, headings[tag])
		}
	}

	if hasLogin {
		t.Errorf("did not expect login form to be detected")
	}
}
