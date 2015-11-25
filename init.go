package binmsg

import (
	"errors"
	"strconv"
	"time"
)

var referenceTime time.Time

func init() {
	_ = initBinMsg()
}

func initBinMsg() error {
	//Below:
	//the error from time.Parse is igonored, because the the value string:
	//"2305-01-01T00:00:00Z" and the layout string: time.RFC3339 does not
	//trigger an error (the error is always nil) ...
	referenceTime, _ = time.Parse(time.RFC3339, "2305-01-01T00:00:00Z")

	referenceTime = referenceTime.UTC()

	ErrPayloadSizeTooSmall = errors.New(
		"The payload size is less than " + strconv.Itoa(PayloadOctets) +
			" bytes long.")
	return nil
}
