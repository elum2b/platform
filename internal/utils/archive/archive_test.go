package archive

import (
	"testing"
	"time"
)

func TestFileName(t *testing.T) {
	now := time.Date(2026, time.February, 1, 12, 0, 0, 123000000, time.UTC)

	tests := map[string]string{
		"calendar":  "calendar_1769947200123.zip",
		"reference": "reference_1769947200123.zip",
	}

	for service, want := range tests {
		if got := fileName(service, now); got != want {
			t.Errorf("fileName(%q) = %q, want %q", service, got, want)
		}
	}
}
