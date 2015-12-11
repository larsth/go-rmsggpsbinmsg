package binmsg

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestInitBinMsg(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	var errString = "The payload size is less than " +
		strconv.Itoa(PayloadOctets) +
		" bytes long."
	var tid, _ = time.Parse(time.RFC3339, "2305-01-01T00:00:00Z")

	initBinMsg()
	if referenceTime.IsZero() {
		t.Fatal("'referenceTime' is zero (0), that is, not initialized.")
	}
	if referenceTime.Equal(tid) == false {
		t.Fatal("'referenceTime's value is not equal to \"2305-01-01T00:00:00Z\".")
	}
	if ErrPayloadSizeTooSmall == nil {
		t.Fatal("'ErrRawDatagramSizeTooSmall' is <nil>")
	}
	if strings.Compare(errString, ErrPayloadSizeTooSmall.Error()) != 0 {
		t.Fatal("'ErrRawDatagramSizeTooSmall' is not equal to\"" + errString + "\"")
	}
}
