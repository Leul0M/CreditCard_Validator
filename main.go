package main

import (
	"fmt"
)

func main() {
	fmt.Println("Credit Card Validator")
	fmt.Println("Enter card numbers to check (type 'exit' to quit)")

	for {
		fmt.Print("Enter card number: ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		if ValidString(input) {
			fmt.Printf("✅ %s is valid\n\n", input)
		} else {
			fmt.Printf("❌ %s is invalid\n\n", input)
		}
	}
}
