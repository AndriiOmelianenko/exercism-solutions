// Package clock implements clock functionality
package clock

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	hour, minute int
}

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

func (c Clock) Add(minutes int) Clock {
	c.minute = c.minute + minutes
	hoursFromMinutes := c.minute / 60
	c.minute = c.minute % 60
	c.hour = (c.hour + hoursFromMinutes) % 24
	if c.minute < 0 {
		c.minute = 60 + c.minute
		c.hour = c.hour - 1
	}
	if c.hour < 0 {
		c.hour = 24 + c.hour
	}
	return Clock{c.hour, c.minute}
}
