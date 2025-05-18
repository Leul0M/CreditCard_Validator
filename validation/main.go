package main

import (
	"fmt"
)

// checks the number is valid or not based on Luhn algorithm
func main() {
	number := 00000000 // Example number
	if Valid(number) {
		fmt.Println("The Number:", number, "is a valid number ✅")
	} else {
		fmt.Println("The Number:", number, "is not a valid number ❌")
	}
}

// CalculateLuhn calculates the check number based on Luhn algorithm
// It returns 0 if the number is already valid
func CalculateLuhn(number int) int {
	checkNumber := checksum(number)

	if checkNumber == 0 {
		return 0
	}
	return 10 - checkNumber
}

// "Valid" check number is valid or not based on Luhn algorithm
func Valid(number int) bool {
	if number < 0 {
		return false
	}
	return (number%10+checksum(number/10))%10 == 0
}

// checksum calculates the Luhn checksum of a number
// It returns the checksum digit
func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}

		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
