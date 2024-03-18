package ui

import (
	"fmt"
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

		return input
	}
}
