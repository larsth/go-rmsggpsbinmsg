package testdata

import "time"

//Time returns a time.Time equal to "2015-11-21T8:41:55Z".
func Time() time.Time {
	return time.Date(2015, 11, 21, 8, 41, 55, 0, time.UTC)
}
