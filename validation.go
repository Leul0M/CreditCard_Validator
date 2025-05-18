package main

import (
	"strings"
	"unicode"
)

// ValidString checks if a credit card number is valid using Luhn algorithm
func ValidString(cardNumber string) bool {
	cleaned := sanitizeInput(cardNumber)
	if len(cleaned) < 2 {
		return false
	}

	luhnSum := 0
	for i, r := range cleaned {
		digit := int(r - '0')
		if (len(cleaned)-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}
		luhnSum += digit
	}
	return luhnSum%10 == 0
}

func sanitizeInput(cardNumber string) string {
	var builder strings.Builder
	for _, r := range cardNumber {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
