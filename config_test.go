package eslog

import (
	"testing"
)

func TestFormat_String(t *testing.T) {
	tests := []struct {
		name     string
		format   Format
		expected string
	}{
		{"JSON format", JSONFormat, "json"},
		{"TEXT format", TextFormat, "text"},
		{"Unknown format", Format(99), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.format.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestParseFormat(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    Format
		shouldError bool
	}{
		{"json lowercase", "json", JSONFormat, false},
		{"text lowercase", "text", TextFormat, false},
		{"invalid format", "invalid", TextFormat, true},
		{"empty string", "", TextFormat, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseFormat(tt.input)
			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}
