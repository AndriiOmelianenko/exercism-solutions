// Package clock implements clock functionality
package clock

import (
	"fmt"
)

const testVersion = 4

// Clock has hour and minute
type Clock struct {
	hour, minute int
}

// New creates new clock
func New(hour, minute int) Clock {
	hoursFromMinutes := minute / 60
	minute = minute % 60
	hour = (hour + hoursFromMinutes) % 24
	if minute < 0 {
		minute = 60 + minute
		hour = hour - 1
	}
	if hour < 0 {
		hour = 24 + hour
	}
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.hour, c.minute)
}

// Add adds minutes to your clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract adds minutes to your clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
