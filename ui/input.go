package ui

import (
	"fmt"
	"github.com/NietThijmen/ShoppingCart/validator"
)

func Input(title string, length int) string {
	fmt.Println(title)
	for {
		var input string

		_, err := fmt.Scan(&input)
		if err != nil {
			return ""
		}

		if len(input) > length {
			fmt.Println("Input too long")
			continue
		}

		// Fixes bug issue: 1
		if validator.ContainsIllegalCharacter(input) {
			fmt.Println("Input contains illegal character")
			continue
		}

		return input
	}
}
