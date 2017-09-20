package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds 1e18 nanoseconds (1e9 seconds)
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e18)
}
