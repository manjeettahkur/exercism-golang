package triangle

import (
	//"fmt"
	"math"
	"sort"
)

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

const testVersion = 3

// Kind defines an integer pointing to which kind of triangule we get
type Kind int

// KindFromSides determine which can of triangule is the three dots it recieves
// as an argument.
func KindFromSides(a, b, c float64) Kind {

	t := []float64{a, b, c}
	sort.Float64s(t)

	for _, value := range t {
		if value <= 0 || math.IsNaN(value) || math.IsInf(value, 0) {
			return NaT
		}
	}
	if t[2] > t[1]+t[0] {
		return NaT
	}
	if t[0] == t[1] && t[1] == t[2] {
		return Equ
	}
	if t[0] == t[1] || t[1] == t[2] {
		return Iso
	}
	if t[0] != t[1] && t[1] != t[2] && t[2] != t[0] {
		return Sca
	}
	return NaT
}
