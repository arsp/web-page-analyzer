package validator

import "testing"

// TestValidateURL uses table driven tests,
// which is the preferred testing style in Go.
func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid https url",
			input:   "https://example.com",
			wantErr: false,
		},
		{
			name:    "valid http url",
			input:   "http://example.com",
			wantErr: false,
		},
		{
			name:    "empty url",
			input:   "",
			wantErr: true,
		},
		{
			name:    "missing scheme",
			input:   "example.com",
			wantErr: true,
		},
		{
			name:    "unsupported scheme",
			input:   "ftp://example.com",
			wantErr: true,
		},
		{
			name:    "malformed url",
			input:   "http://",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateURL(tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("did not expect error, got %v", err)
			}
		})
	}
}
