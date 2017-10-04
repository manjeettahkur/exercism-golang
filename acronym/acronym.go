package acronym

import (
	"fmt"
	"strings"
)

const testVersion = 3

// Abbreviate creates acronym in base of input names
func Abbreviate(i string) string {
	phrase := strings.FieldsFunc(i, split)
	var acronym string
	for _, word := range phrase {
		acronym = fmt.Sprintf("%s%s", acronym, strings.ToUpper(string(word[0])))
	}
	return acronym
}

// I didn't find a way to split in base of multiple delimiters
func split(r rune) bool {
	return r == '-' || r == ' '
}
