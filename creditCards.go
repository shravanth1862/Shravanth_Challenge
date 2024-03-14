package main

import (
	"fmt"
)

func validateCreditCard(cardNumber string) string {

	if cardNumber[0] != '4' && cardNumber[0] != '5' && cardNumber[0] != '6' { //make sure card starts with either 4 or 5 or 6
		return "Invalid1"
	}

	length := len(cardNumber)
	containsHyphens := false
	if length == 19 { //makes sure length is 16 exactly. only allow greter than 16 if hyphens are present at expected indicies
		if cardNumber[4] != '-' && cardNumber[9] != '-' && cardNumber[14] != '-' {
			return "Invalid2"
		} else {
			containsHyphens = true
		}
	} else if length < 16 {
		return "Invalid3"
	} else if length > 16 {
		return "Invalid4"
	}

	for i := 0; i < len(cardNumber); i++ {
		if cardNumber[i] < '0' || cardNumber[i] > '9' { //check if card has any number other than a digit
			if containsHyphens == true && (i == 4 || i == 9 || i == 14) { //any number other than a digit with the exception of hyphens at specific indexes
				continue
			}
			return "Invalid5"
		}
	}

	// Check if the card number contains 4 or more of same digit
	for i := 0; i < len(cardNumber)-3; i++ {
		if cardNumber[i] == cardNumber[i+1] && cardNumber[i] == cardNumber[i+2] && cardNumber[i] == cardNumber[i+3] {
			return "Invalid6"
		}
	}

	return "Valid"
}

func main() {

	var input string

	for {
		fmt.Println("Enter a card number to validate or type 'end' to exit.")
		fmt.Scanln(&input)
		if input == "end" {
			break
		}
		fmt.Println(validateCreditCard(input))
	}

	fmt.Println("Program ended")
}
