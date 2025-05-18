package main

import "testing"

func TestValidString(t *testing.T) {
	tests := []struct {
		name     string
		cardNum  string
		expected bool
	}{
		{"Valid Visa", "4539 1488 0343 6467", true},
		{"Invalid Empty", "", false},
		{"Invalid Non-Digits", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidString(tt.cardNum); got != tt.expected {
				t.Errorf("ValidString(%q) = %v, want %v", tt.cardNum, got, tt.expected)
			}
		})
	}
}
