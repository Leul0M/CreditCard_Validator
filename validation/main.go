package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	number := 79927398713
	numberStr := fmt.Sprintf("%d", number)
	if ValidString(numberStr) {
		fmt.Println(number, "is a valid number")
	} else {
		fmt.Println(number, "is not a valid number")
	}
}

// sanitizeInput removes all non-digit characters from the string.
func sanitizeInput(cardNumber string) string {
	var builder strings.Builder
	for _, r := range cardNumber {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// ValidString checks if a credit card number (string) is valid using Luhn's algorithm.
func ValidString(cardNumber string) bool {
	cleaned := sanitizeInput(cardNumber)
	if len(cleaned) < 2 { // Minimum length for Luhn (e.g., "0" is invalid)
		return false
	}

	luhnSum := 0
	for i, r := range cleaned {
		digit := int(r - '0') // Convert rune to int (e.g., '5' → 5)

		// Double every second digit from the right
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
