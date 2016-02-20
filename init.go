package binmsg

import (
	"strconv"
	"time"

	"github.com/juju/errors"
)

var referenceTime time.Time

const refTimeString = "2305-01-01T00:00:00Z"

func init() {
	initBinMsg()
}

func initBinMsg() {
	//the error from time.Parse is not checked , because it does not trigger
	//an error with format string time.RFC3339, and 'refTimeString'
	referenceTime, _ = time.Parse(time.RFC3339, refTimeString)
	referenceTime = referenceTime.UTC()
	ErrPayloadSizeTooSmall = errors.New(
		"The payload size is less than " + strconv.Itoa(PayloadOctets) +
			" bytes long.")
}
