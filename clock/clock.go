package clock

import "fmt"

const testVersion = 4

// You can find more details and hints in the test file.

type Clock struct {
	minutes int
	time    string
}

func New(hour, minute int) Clock {
	var reloj Clock
	reloj.minutes = (hour*60 + minute) % (24 * 60)
	if reloj.minutes < 0 {
		reloj.minutes += (60 * 24)
	}
	return reloj
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

func (c Clock) Add(minutes int) Clock {
	c.minutes = (c.minutes + minutes) % (24 * 60)
	if c.minutes < 0 {
		c.minutes += (60 * 24)
	}
	return c
}

/*
func (Clock) String() string {
}

func (Clock) Add(minutes int) Clock {
}

*/

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
