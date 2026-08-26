package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
    var billionSeconds time.Duration = 1000000000 * time.Second
    t = t.Add(billionSeconds)
	return t
}
