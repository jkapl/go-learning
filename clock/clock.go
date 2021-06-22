package clock

import "math"

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	totalMin := (hour * 60) + minute 
	if totalMin < 0 {

	}
	hours := int(math.Floor(float64(totalMin) / 60))
	mins := int(totalMin % 60)
	
	c := Clock{hours, mins}
	return c
	
}

func (c Clock) String() string {
	return string(c.hour) + ":" + string(c.minute)
}

func (c Clock) Add(minutes int) Clock {
	// if minutes >= 60
	return Clock{1,1}
}

func (c Clock) Subtract(minutes int) Clock {

	return Clock{1,1}
}