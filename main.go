package main

import (
	"fmt"
)

func main() {
	fmt.Println("Credit Card Validator💳")
	fmt.Println("Enter card numbers to check (type 'q' to quit)")

	for {
		fmt.Print("Enter card number: ")
		var input string
		fmt.Scanln(&input)

		if input == "q" {
			break
		}

		if ValidString(input) {
			fmt.Printf("✅ %s is valid\n\n", input)
		} else {
			fmt.Printf("❌ %s is invalid\n\n", input)
		}
	}
}
