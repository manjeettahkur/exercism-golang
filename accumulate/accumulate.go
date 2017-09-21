package accumulate

const testVersion = 1

// Accumulate is a function to transform strings
func Accumulate(s []string, f func(string) string) []string {
	sReturn := []string{}
	for _, word := range s {
		sReturn = append(sReturn, f(word))
	}
	return sReturn
}
