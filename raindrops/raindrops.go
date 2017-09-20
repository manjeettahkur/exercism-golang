package raindrops

import "fmt"

const testVersion = 3

// Convert allow you to convert number to a string
func Convert(n int) string {
	var rValue string
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			switch i {
			case 3:
				rValue += "Pling"
			case 5:
				rValue += "Plang"
			case 7:
				rValue += "Plong"
			}
		}
	}
	if rValue == "" {
		rValue = fmt.Sprintf("%d", n)
	}
	return rValue
}
