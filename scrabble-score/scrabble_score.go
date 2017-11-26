package scrabble

// It doesn't look super clean to me. I am sure is better way, just need to think more :(

import (
	"regexp"
	"strings"
)

func Score(word string) int {
	var value1 = regexp.MustCompile("[AEIOULNRST]")
	var value2 = regexp.MustCompile("[DG]")
	var value3 = regexp.MustCompile("[BCMP]")
	var value4 = regexp.MustCompile("[FHVWY]")
	var value5 = regexp.MustCompile("[K]")
	var value8 = regexp.MustCompile("[JX]")
	var value10 = regexp.MustCompile("[QZ]")

	var score int

	for _, letter := range word {
		switch {
		case value1.MatchString(strings.ToUpper(string(letter))):
			score = score + 1
		case value2.MatchString(strings.ToUpper(string(letter))):
			score = score + 2
		case value3.MatchString(strings.ToUpper(string(letter))):
			score = score + 3
		case value4.MatchString(strings.ToUpper(string(letter))):
			score = score + 4
		case value5.MatchString(strings.ToUpper(string(letter))):
			score = score + 5
		case value8.MatchString(strings.ToUpper(string(letter))):
			score = score + 8
		case value10.MatchString(strings.ToUpper(string(letter))):
			score = score + 10
		}
	}

	return score
}
