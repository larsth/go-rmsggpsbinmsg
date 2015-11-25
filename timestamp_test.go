package binmsg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

type TtdTimeStampMarshal struct {
	Init     func(testItem *TtdTimeStampMarshal)
	TS       TimeStamp
	WantPStr string
	WantP    []byte
}

var tdTimeStampMarshalJSONWant = []*TtdTimeStampMarshal{
	//Test 0:
	&TtdTimeStampMarshal{
		Init: func(testItem *TtdTimeStampMarshal) {
			testItem.TS.Time = time.Date(2015, 01, 23, 12, 34, 59, 0, time.UTC)
			testItem.WantPStr = testItem.TS.Time.Format(`"` + time.RFC3339Nano + `"`)
			testItem.WantP = []byte(testItem.WantPStr)
		},
	},
}

func TestTimeStampMarshalJSON(t *testing.T) {
	var gotStr string
	//(*TimeStamp).MarshalJSON uses the embedded (*time.Time)MarshalJSON() method:
	//Which is (from: $GOROOT/src/time/time.go, lines 930 to 939):
	// 	//MarshalJSON implements the json.Marshaler interface.
	//	// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
	//	func (t Time) MarshalJSON() ([]byte, error) {
	//		if y := t.Year(); y < 0 || y >= 10000 {
	//			// RFC 3339 is clear that years are 4 digits exactly.
	//			// See golang.org/issue/4556#c15 for more discussion.
	//			return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	//		}
	//		return []byte(t.Format(`"` + RFC3339Nano + `"`)), nil
	//	}
	//... , so err is always <nil>, if years is 4 digits:

	log.SetOutput(ioutil.Discard)

	for i, testItem := range tdTimeStampMarshalJSONWant {
		if testItem.Init != nil {
			testItem.Init(testItem)
		}
		gotP, _ := testItem.TS.MarshalJSON()
		if bytes.Compare(gotP, testItem.WantP) != 0 {
			if gotP != nil {
				gotStr = string(gotP)
			} else {
				//gotP == nil is true:
				gotStr = fmt.Sprintf("%v", gotP)
			}
			t.Fatalf("Test %d: Got '%s', but want '%s'\n",
				i, gotStr, testItem.WantPStr)
		}
	}
}

type TtdTimeStampUnmarshal struct {
	Init     func(testItem *TtdTimeStampUnmarshal)
	WantTime time.Time
	Data     string
	WantErr  error
}

var tdTimeStampUnmarshalJSONWant = []*TtdTimeStampUnmarshal{
	//Test 0::
	&TtdTimeStampUnmarshal{
		Init: func(testItem *TtdTimeStampUnmarshal) {
			testItem.WantTime = time.Date(2015, 01, 23, 12, 34, 59, 0, time.UTC)
			testItem.Data = testItem.WantTime.Format(`"` + time.RFC3339Nano + `"`)
			testItem.WantErr = nil
		},
	},
}

func TestTimeStampUnmarshalJSON(t *testing.T) {
	//(*TimeStamp).MarshalJSON uses the embedded (*time.Time)UnmarshalJSON()
	//method, which is implemented like this in $GOROOT/src/time/time.go,
	//lines 930 to 939):
	//	func (t *Time) UnmarshalJSON(data []byte) (err error) {
	// 		Fractional seconds are handled implicitly by Parse.
	//		*t, err = Parse(`"`+RFC3339+`"`, string(data))
	//		return
	//	}
	var (
		data   []byte
		gotTS  TimeStamp
		gotErr error
	)

	log.SetOutput(ioutil.Discard)

	for i, testItem := range tdTimeStampUnmarshalJSONWant {
		if testItem.Init != nil {
			testItem.Init(testItem)
		}
		data = []byte(testItem.Data)
		gotErr = gotTS.UnmarshalJSON(data)

		if gotTS.Time.Equal(testItem.WantTime) == false {
			t.Fatalf("Test: %d:: Got: '%s', Want: '%s'\n",
				i, gotTS.Time.String(), testItem.WantTime.String())
		}
		if s, ok := errorCheck(i, gotErr, testItem.WantErr); ok == false {
			t.Fatal(s)
		}
	}
}

//type TtdTimeStampMarshalBinary struct {
//	Init     func(testItem *TtdTimeStampMarshalBinary)
//	TS       TimeStamp
//	WantData []byte
//	WantErr  error
//}

//var tdTimeStampMarshalBinary = []*TtdTimeStampMarshalBinary{
//	&TtdTimeStampMarshalBinary{
//		Init: func(testItem *TtdTimeStampMarshalBinary) {
//			testItem.TS.Time = time.Date(2015, 01, 23, 12, 34, 59, 0, time.UTC)
//		},
//		WantData: []byte{129, 6, 76, 30, 13, 87, 190, 0},
//		WantErr:  nil,
//	},
//}

func TestTimeStampMarshalBinary(t *testing.T) {
	const (
		year, month, day      = 2015, 01, 23
		hour, minute, seconds = 12, 34, 59
		nanoseconds           = 0
	)
	var (
		ts             *TimeStamp
		wantData       = []byte{129, 6, 76, 30, 13, 87, 190, 0}
		v1, v2, v3, v4 byte
		v5, v6, v7, v8 byte
		gotData        []byte
	)

	ts = new(TimeStamp)
	ts.Time = time.Date(year, month, day, hour, minute, seconds, nanoseconds, time.UTC)
	gotData = make([]byte, 8)
	log.SetOutput(ioutil.Discard)

	v1, v2, v3, v4, v5, v6, v7, v8 = ts.marshalBytes()
	gotData[0] = v1
	gotData[1] = v2
	gotData[2] = v3
	gotData[3] = v4
	gotData[4] = v5
	gotData[5] = v6
	gotData[6] = v7
	gotData[7] = v8

	if s, ok := byteSliceCheck(0, gotData, wantData); ok == false {
		t.Fatal(s)
	}
}

func TestTimeStampUnmarshalBytes(t *testing.T) {
	const (
		year, month, day      = 2015, 01, 23
		hour, minute, seconds = 12, 34, 59
		nanoseconds           = 0
	)
	var (
		ts             *TimeStamp
		v1, v2, v3, v4 = byte(129), byte(6), byte(76), byte(30)
		v5, v6, v7, v8 = byte(13), byte(87), byte(190), byte(0)
	)
	ts = new(TimeStamp)
	ts.Time = time.Date(year, month, day, hour, minute, seconds, nanoseconds, time.UTC)

	log.SetOutput(ioutil.Discard)

	ts.unmarshalBytes(v1, v2, v3, v4, v5, v6, v7, v8)
	if ts.Time.Year() != year {
		t.Fatal(mkIntErrString("Year", ts.Time.Year(), year))
	}
	if ts.Time.Month() != month {
		s := mkStrErrString("Month",
			ts.Time.Month().String(),
			(time.Month(month)).String())
		t.Fatal(s)
	}
	if ts.Time.Day() != day {
		t.Fatal(mkIntErrString("Day", ts.Time.Day(), day))
	}
	if ts.Time.Hour() != hour {
		t.Fatal(mkIntErrString("Hour", ts.Time.Hour(), hour))
	}
	if ts.Time.Minute() != minute {
		t.Fatal(mkIntErrString("Minute", ts.Time.Minute(), minute))
	}
	if ts.Time.Second() != seconds {
		t.Fatal(mkIntErrString("Second", ts.Time.Second(), seconds))
	}
	if ts.Time.Nanosecond() != 0 {
		t.Fatal(mkIntErrString("Nanosecond", ts.Time.Nanosecond(), nanoseconds))
	}
}

func BenchmarkTimeStampMarshalJSON(b *testing.B) {
	var ts TimeStamp
	ts.Time = time.Date(2015, 12, 11, 24, 31, 9, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ts.MarshalJSON()
	}
}

func BenchmarkTimeStampUnmarshalJSON(b *testing.B) {
	var ts TimeStamp
	var data []byte
	ts.Time = time.Date(2015, 12, 11, 24, 31, 9, 0, time.UTC)
	data, _ = ts.MarshalJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ts.UnmarshalJSON(data)
	}
}

func BenchmarkTimeStampMarshalBytes(b *testing.B) {
	var ts TimeStamp
	ts.Time = time.Date(2015, 12, 11, 24, 31, 9, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _, _, _, _, _, _ = ts.marshalBytes()
	}
}

func BenchmarkTimeStampUnmarshalBytes(b *testing.B) {
	var ts TimeStamp
	var v1, v2, v3, v4, v5, v6, v7, v8 byte
	v1, v2, v3, v4 = 0x01, 0x23, 0x45, 0x67
	v5, v6, v7, v8 = 0x89, 0xab, 0xcd, 0xef
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ts.unmarshalBytes(v1, v2, v3, v4, v5, v6, v7, v8)
	}
}
