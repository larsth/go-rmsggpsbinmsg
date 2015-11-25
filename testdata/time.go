package testdata

import "time"

func Time() time.Time {
	return time.Date(2015, 11, 21, 8, 41, 55, 0, time.UTC)
}
