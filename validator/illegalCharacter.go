package validator

import "strings"

var illegalCharacters = []string{
	";",
	"&",
}

func ContainsIllegalCharacter(value string) bool {
	for _, illegalCharacter := range illegalCharacters {
		if strings.Contains(value, illegalCharacter) {
			return true
		}
	}

	return false
}
