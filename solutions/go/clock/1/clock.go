package clock

import "fmt"

type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	total := h*60 + m

	// Normalize to a 24-hour clock.
	total = ((total % 1440) + 1440) % 1440

	return Clock{
		h: total / 60,
		m: total % 60,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}