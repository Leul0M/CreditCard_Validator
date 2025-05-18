package main

import "testing"

// TestValidString tests the ValidString function
// for now it workes 🤷‍♂️
func TestValidString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Visa", "4539 1488 0343 6467", true},
		{"Valid Mastercard", "5555 5555 5555 4444", true},
		{"Invalid Number", "1234 5678 9012 3456", false},
		{"Empty Input", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidString(tt.input); got != tt.expected {
				t.Errorf("ValidString() = %v, want %v", got, tt.expected)
			}
		})
	}
}
