package binmsg

//FixMode is a type used for indication no GPS fix, 2D GPS fix, and 3D GPS fix.
type FixMode byte

func (f *FixMode) String() string {
	switch *f {
	case FixNotSeen:
		return "Not Seen"
	case FixNone:
		return "None"
	case Fix2D:
		return "2D"
	case Fix3D:
		return "3D"
	}
	return "Unknown FixMode value" //make compiler happy
}

//MarshalJSON implements the json.Marshaler interface.
//It marshals a valid FixMode value to the ASCII 8-bit/UTF-8 values:
//"0", "1", "2", or "3".
//If the FixMode is valid it returns that value in a 1-byte byte slice, and
//a nil error.
//
//If the FixMode is not a valid FixMode value, then the
//nil byte slice, and the ErrUnknownFixMode error are returned.
func (f *FixMode) MarshalJSON() ([]byte, error) {
	var (
		p = make([]byte, 1, 1)
	)
	switch *f {
	case FixNotSeen:
		p[0] = 0x30 //ASCII/UTF8: "0"
		return p, nil
	case FixNone:
		p[0] = 0x31 //ASCII/UTF8: "1"
		return p, nil
	case Fix2D:
		p[0] = 0x32 //ASCII/UTF8: "2"
		return p, nil
	case Fix3D:
		p[0] = 0x33 //ASCII/UTF8: "3"
		return p, nil
	}
	return nil, ErrUnknownFixMode
}

//UnmarshalJSON implements the json.Unmarshaler interface.
//It unmarshals the ASCII 8-bit/UTF-8 values: "0", "1", "2", or "3" to a
//FixMode value. If any other value than the above vlauea are met, then
//UnmarshalJSON return the ErrUnknownFixMode, and sets the FixMode value to 253.
func (f *FixMode) UnmarshalJSON(p []byte) error {
	if len(p) == 0 {
		(*f) = FixMode(byte(252))
		return ErrNilSlice
	}
	switch p[0] {
	case 0x30: //ASCII/UTF8: "0"
		(*f) = FixNotSeen
		return nil
	case 0x31: //ASCII/UTF8: "1"
		(*f) = FixNone
		return nil
	case 0x32: //ASCII/UTF8: "2"
		(*f) = Fix2D
		return nil
	case 0x33: //ASCII/UTF8: "3"
		(*f) = Fix3D
		return nil
	}
	(*f) = FixMode(byte(253))
	return ErrUnknownFixMode
}

func (f *FixMode) marshalByte() (v byte, err error) {
	err = nil
	switch *f {
	case FixNotSeen:
		v = 0
		return v, nil
	case FixNone:
		v = 1
		return v, nil
	case Fix2D:
		v = 2
		return v, nil
	case Fix3D:
		v = 3
		return v, nil
	}
	err = ErrUnknownFixMode
	v = 255
	return
}

func (f *FixMode) unmarshalByte(data byte) error {
	switch data {
	case 0:
		(*f) = FixNotSeen
		return nil
	case 1:
		(*f) = FixNone
		return nil
	case 2:
		(*f) = Fix2D
		return nil
	case 3:
		(*f) = Fix3D
		return nil
	}
	(*f) = FixMode(byte(255))
	return ErrUnknownFixMode
}

const (
	//FixNotSeen means that there is no knowledge of what kind of fix a GPS has.
	FixNotSeen FixMode = 0
	//FixNone means that the GPS hasn´t a fix.
	FixNone FixMode = 1
	//Fix2D means that the GPS has a 2D fix.
	Fix2D FixMode = 2
	//Fix3D means that the GPS has a 3D fix.
	Fix3D FixMode = 3
)
