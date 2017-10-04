package bob

import (
	"regexp"
)

const testVersion = 3

// Hey is a robot who replies in different ways in base of what you say.
func Hey(i string) string {

	var questionMark = regexp.MustCompile(".*\\?\\s*$")
	var questionItself = regexp.MustCompile("^[A-Z ]+\\?$")
	var yellMark = regexp.MustCompile("(!$|[A-Z]!?$)")
	var yellItself = regexp.MustCompile("[ a-z]!$")

	var silence = regexp.MustCompile("(^[^\\.!0-9a-z]+$|^$)")

	switch {
	case questionMark.MatchString(i):
		if questionItself.MatchString(i) {
			return "Whoa, chill out!"
		}
		return "Sure."
	case yellMark.MatchString(i):
		if yellItself.MatchString(i) {
			return "Whatever."
		}
		return "Whoa, chill out!"
	case silence.MatchString(i):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
