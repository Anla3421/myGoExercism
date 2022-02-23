package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour int
	min  int
}

func New(h, m int) Clock {
	h, m = timeFormatCheck(h, m)
	return Clock{
		hour: h,
		min:  m,
	}
}

func (c Clock) Add(m int) Clock {
	c.min += m
	c.hour, c.min = timeFormatCheck(c.hour, c.min)
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.min -= m
	c.hour, c.min = timeFormatCheck(c.hour, c.min)
	return c
}

func (c Clock) String() string {
	myTime := fmt.Sprintf("%.2d:%.2d", c.hour, c.min)
	return myTime
}

func timeFormatCheck(h, m int) (int, int) {
	leave := false
	for {
		switch {
		case m >= 60:
			h += 1
			m -= 60
		case h >= 24:
			h -= 24
		case m < 0:
			h -= 1
			m += 60
		case h < 0:
			h += 24
		default:
			leave = true
		}
		if leave == true {
			break
		}
	}
	return h, m
}
