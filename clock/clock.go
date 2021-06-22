package clock

import (
	"math"
	"strconv"
)

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	totalMin := (hour * 60) + minute 
	hours := int(math.Floor(float64(totalMin) / 60))
	mins := int(totalMin % 60)
	
	if totalMin < 0 {
		finalHour := 24 - hours;
		finalMin := 60 - mins;
		return Clock{finalHour, finalMin}
	}

	return Clock{hours, mins}
}

func (c Clock) String() string {
	var hour string
	var minute string
	if (c.hour < 10) {
		hour = "0" + strconv.Itoa(c.hour)
	} else {
		hour = strconv.Itoa(c.hour)
	}
	if (c.minute < 10) {
		minute = "0" + strconv.Itoa(c.minute)
	} else {
		minute = strconv.Itoa(c.minute)
	}
	return hour + ":" + minute
	
}

func (c Clock) Add(minutes int) Clock {
	hours := int(c.hour)
	minute := int(c.minute)
	newMinutes := minutes + minute
	return Clock{hours,newMinutes}
}

func (c Clock) Subtract(minutes int) Clock {
	hours := int(c.hour)
	minute := int(c.minute)
	newMinutes := minute - minutes
	return Clock{hours,newMinutes}
}