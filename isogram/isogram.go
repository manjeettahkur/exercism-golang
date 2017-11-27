package isogram

import (
	"strings"
)

// IsIsogram checks if a word is a Isogram
func IsIsogram(word string) bool {
	if word == "" {
		return true
	}

	duplicates := make(map[string]int)

	for _, l := range word {
		letra := strings.ToUpper(string(l))
		if _, ok := duplicates[letra]; ok {
			if letra == "-" || letra == " " {
				continue
			}
			return false
		} else {
			duplicates[letra] = 1
		}
	}
	return true
}
