// Package gigasecond calculates the moment when someone has lived for 10^9 seconds
package gigasecond

import (
	"math"
	"time"
)

const testVersion = 4

func AddGigasecond(myDate time.Time) time.Time {
	return myDate.Add(time.Duration(math.Pow10(18)))
}
