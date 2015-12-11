package binmsg

import (
	"encoding/binary"
	"time"
)

const timeStampOctets = 8

//TimeStamp contains the timestamp, which is used to distinghish between an old
//message and a new message.
type TimeStamp struct {
	Time time.Time
}

//MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision
//added if present.
func (t *TimeStamp) MarshalJSON() ([]byte, error) {
	return t.Time.MarshalJSON()
}

//UnmarshalJSON implements the json.Unmarshaler interface.
//The time is expected to be a quoted string in RFC 3339 format.
func (t *TimeStamp) UnmarshalJSON(data []byte) error {
	return t.Time.UnmarshalJSON(data)
}

func (t *TimeStamp) marshalBytes() (v1, v2, v3, v4, v5, v6, v7, v8 byte) {
	var (
		d   time.Duration
		i64 int64
		p   = make([]byte, 8)
	)

	d = t.Time.Sub(referenceTime)
	i64 = int64(d)
	binary.BigEndian.PutUint64(p, uint64(i64))

	v1 = p[0]
	v2 = p[1]
	v3 = p[2]
	v4 = p[3]
	v5 = p[4]
	v6 = p[5]
	v7 = p[6]
	v8 = p[7]
	return
}

func (t *TimeStamp) unmarshalBytes(v1, v2, v3, v4, v5, v6, v7, v8 byte) {
	var (
		p   = make([]byte, 8)
		i64 int64
		d   time.Duration
	)

	p[0] = v1
	p[1] = v2
	p[2] = v3
	p[3] = v4
	p[4] = v5
	p[5] = v6
	p[6] = v7
	p[7] = v8

	i64 = int64(binary.BigEndian.Uint64(p))
	d = time.Duration(i64)
	t.Time = referenceTime.Add(d)
}
