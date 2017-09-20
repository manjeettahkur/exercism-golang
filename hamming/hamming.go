package hamming

import "errors"

const testVersion = 6

// Distance get the Hamming distance for two strings of equal lenght
func Distance(a, b string) (int, error) {
	var HammingDistance int
	if len(a) != len(b) {
		return -1, errors.New("The strings have different lenght")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			HammingDistance++
		}
	}
	return HammingDistance, nil
}
