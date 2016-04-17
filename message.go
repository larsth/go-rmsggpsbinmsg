package binmsg

import (
	"strconv"
	"time"

	"github.com/juju/errors"
)

const messageOctets = timeStampOctets + gpsOctets

//Message is a type that contains a TimeStamp type (when was this message
//created?), and a Gps type.
type Message struct {
	//TimeStamp octets: timeStampOctets(=8) bytes (type time.Duration is an int64 value)
	TimeStamp TimeStamp `json:"timestamp"`
	//Gps octets: gpsOctet bytes
	Gps Gps `json:"gps"`
}

//MarshalBinary marshals the struct fields from type Message into a
//binary representation of type Message, which are returned in a byte slice.
func (m *Message) MarshalBinary() ([]byte, error) {
	var (
		v1, v2, v3, v4, v5, v6, v7, v8 byte
		p                              = make([]byte, messageOctets)
		err                            error
	)

	//Struct fields from type Message ...
	//	1.0 Marshal The Gps structure to binary ...
	//	1.1 Marshal the FixMode value
	p[0], err = m.Gps.FixMode.MarshalByte()
	if err != nil {
		annotatedErr := errors.Annotate(err,
			"Payload.Message.Gps.FixMode.MarshalByte() error")
		return nil, annotatedErr
	}
	//Example FixMode value: Fix3D -> 0x03
	//
	// 1.2 Marshal the Latitude IEEE 754 32-bit floating-point value
	v1, v2, v3, v4 = float32MarshalBinaryValues(m.Gps.Latitude)
	//Example latitude: float32(55.69147) -> v1=0x42, v2=0x5e, v3=0xc4, v4=0x11
	p[1] = v1
	p[2] = v2
	p[3] = v3
	p[4] = v4
	// 1.3 Marshal the Longitude IEEE 754 32-bit floating-point value
	v1, v2, v3, v4 = float32MarshalBinaryValues(m.Gps.Longitude)
	//Example longitude: float32(12.61681) -> v1 = 0x41, v2=0x49, v3=0xde, v4=0x74
	p[5] = v1
	p[6] = v2
	p[7] = v3
	p[8] = v4
	// 1.4 Marshal the Altitude IEEE 754 32-bit floating-point value
	v1, v2, v3, v4 = float32MarshalBinaryValues(m.Gps.Altitude)
	//Example altitude: float32(2.01) -> v1=0x40, v2=0x00, v3=0xa3, v4=0xd7
	p[9] = v1
	p[10] = v2
	p[11] = v3
	p[12] = v4
	// 2.0 Marshal the TimeStamp
	v1, v2, v3, v4, v5, v6, v7, v8 = m.TimeStamp.marshalBytes()
	//Example timestamp: "2015-11-21T08:41:55Z"
	//	and reference time is "2305-01-01T00:00:00Z" ->
	// 		v1=0x81, v2=0x62, v3=0xf2, v4=0xa9,
	// 		v5=0x91, v6=0x2f, v7=0x7e, v8=0x00
	p[13] = v1
	p[14] = v2
	p[15] = v3
	p[16] = v4
	p[17] = v5
	p[18] = v6
	p[19] = v7
	p[20] = v8

	return p, nil
}

//UnmarshalBinary unmarshals a binary representation of a Payload in a byte
//slice into the data a Payload type contains.
func (m *Message) UnmarshalBinary(data []byte) error {
	var (
		f   float32
		err error
		b   byte
		s   []byte
	)

	// 1.0 Data argument octet size check ...
	if len(data) < PayloadOctets {
		return errors.Trace(ErrPayloadSizeTooSmall)
	}

	// 2.0 Unmarshal the GPS POI ...
	// 2.1 Unmarshal the FixMode ...
	b = data[0]
	if err = (m.Gps.FixMode).UnmarshalByte(b); err != nil {
		return errors.Annotate(err,
			"Package.Message.Gps.FixMode).UnmarshalByte error")
	}
	// 2.2 Unmarshal the latitude ...
	s = data[1:5]
	f = float32UnmarshalBinary(s)

	m.Gps.Latitude = f
	// 2.3 Unmarshal the longitude ...
	s = data[5:9]
	f = float32UnmarshalBinary(s)
	m.Gps.Longitude = f
	// 2.4 Unmarshal the altitude ...
	s = data[9:13]
	f = float32UnmarshalBinary(s)
	m.Gps.Altitude = f

	// 2.5 Unmarshal the timestamp
	s = data[13:]
	m.TimeStamp.unmarshalBytes(
		s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7])

	return nil
}

func (m *Message) Strings() (fixmode, alt, lat, lon, timestamp string) {
	fixmode = m.Gps.FixMode.String()
	alt = strconv.FormatFloat(m.Gps.Alt(), 'f', -1, 32)
	lat = strconv.FormatFloat(m.Gps.Lat(), 'f', -1, 32)
	lon = strconv.FormatFloat(m.Gps.Lon(), 'f', -1, 32)
	timestamp = m.TimeStamp.Time.Format(time.RFC3339)
	return
}
