package binmsg

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
	"sync"

	"github.com/juju/errors"
)

const shaSize = sha256.Size224

//PayloadOctets is the amount of octets (bytes) the 'data' byte slice
//returned from (*Payload).MarshalBinary() (data []byte, err error) has.
const PayloadOctets = messageOctets + shaSize

//Secrets is a type that contains the shared secrets, while a Payload had
//been initialized.
type Secrets struct {
	HMACKey []byte //shared HMAC secret key
	Salt    []byte //crypto.Rand generated garbage - lots of it, shared secret
}

//Payload is a type which is a representaion of the payload transmitted
//between RMSG.dk programs.
type Payload struct {
	mutex         sync.Mutex `json:"-"`
	Secrets       Secrets    `json:"-"`
	Message       Message
	messageOctets []byte `json:"-"`
	hMACOctets    []byte `json:"-"`
}

func createHMAC(hmacKey, salt, messageOctets []byte) (octets []byte, err error) {
	var (
		mac hash.Hash
	)

	octets = nil
	if len(hmacKey) == 0 {
		err = errors.Trace(ErrNilHMACSecretKeySlice)
		return
	}
	if len(salt) == 0 {
		err = errors.Trace(ErrNiSaltSlice)
		return
	}
	if messageOctets == nil {
		err = errors.Trace(ErrNilMessageOctetsSlice)
		return
	}
	err = nil

	mac = hmac.New(sha256.New224, hmacKey)

	mac.Write(hmacKey)
	mac.Write(messageOctets)
	mac.Write(salt)

	octets = mac.Sum(nil)
	return
}

//checkHMAC reports whether p.HMACOctets is a valid HMAC tag for p.MessageOctets
func checkHMAC(hmacKey, salt, messageOctets, hmacOctets []byte) (err error) {
	var (
		expectedMAC []byte
		ok          bool
	)

	if len(hmacOctets) != shaSize {
		err = errors.Trace(ErrHMACOctetsWrongSize)
		ok = false
		return
	}
	if expectedMAC, err = createHMAC(hmacKey, salt, messageOctets); err != nil {
		return errors.Annotate(err, "expectedMAC: createHMAC error")
	}

	ok = hmac.Equal(hmacOctets, expectedMAC)
	if ok == false {
		return errors.Trace(ErrHMACcheckFailed)
	}
	return nil
}

//Init initializes a Payload type with the given 'hmacKey' and 'salt' slices.
//Init does a simple zero length/nil check on both given slices, and one of
//them has the zero length or is nil, then an error is reutned, if there are
//no errors, then the nil error is returned.
func (p *Payload) Init(hmacKey []byte, salt []byte) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(hmacKey) == 0 {
		return errors.Trace(ErrNilHMACSecretKeySlice)
	}
	if len(salt) == 0 {
		return errors.Trace(ErrNiSaltSlice)
	}
	p.Secrets.HMACKey = hmacKey
	p.Secrets.Salt = salt
	return nil
}

//New creates a new Payload with the given 'hmacKey' and 'salt' byte slices.
//If ('Payload).Init(hmacKey, salt) returns an error, then a nil Payload and
//the error are returned.
func New(hmacKey []byte, salt []byte) (*Payload, error) {
	p := new(Payload)
	if err := p.Init(hmacKey, salt); err != nil {
		annotatedErr := errors.Annotate(err,
			"Init(hmacKey, salt) error")
		return nil, annotatedErr
	}
	return p, nil
}

//MarshalBinary marshals the data a payload type contains into a binary
//representation of a Payload, which are stored in a byte slice.
func (p *Payload) MarshalBinary() (data []byte, err error) {
	var v1, v2, v3, v4, v5, v6, v7, v8 byte

	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.messageOctets = nil
	data = nil
	if len(p.Secrets.HMACKey) == 0 {
		//HMAC secret is not set ...
		err = errors.Trace(ErrNilHMACSecretKeySlice)
		return
	}
	if len(p.Secrets.Salt) == 0 {
		//Salt (crypto random numbers) are not set ...
		err = errors.Trace(ErrNiSaltSlice)
		return
	}
	p.messageOctets = make([]byte, messageOctets, PayloadOctets)

	//the message ...
	//	1.0 Marshal The Gps structure to binary ...
	//	1.1 Marshal the FixMode value
	p.messageOctets[0], err = p.Message.Gps.FixMode.MarshalByte()
	if err != nil {
		annotatedErr := errors.Annotate(err,
			"Payload.Message.Gps.FixMode.MarshalByte() error")
		return nil, annotatedErr
	}
	//Example FixMode value: Fix3D -> 0x03
	//
	// 1.2 Marshal the Latitude IEEE 754 32-bit floating-point value
	v1, v2, v3, v4 = float32MarshalBinaryValues(p.Message.Gps.Latitude)
	//Example latitude: float32(55.69147) -> v1=0x42, v2=0x5e, v3=0xc4, v4=0x11
	p.messageOctets[1] = v1
	p.messageOctets[2] = v2
	p.messageOctets[3] = v3
	p.messageOctets[4] = v4
	// 1.3 Marshal the Longitude IEEE 754 32-bit floating-point value
	v1, v2, v3, v4 = float32MarshalBinaryValues(p.Message.Gps.Longitude)
	//Example longitude: float32(12.61681) -> v1 = 0x41, v2=0x49, v3=0xde, v4=0x74
	p.messageOctets[5] = v1
	p.messageOctets[6] = v2
	p.messageOctets[7] = v3
	p.messageOctets[8] = v4
	// 1.4 Marshal the Altitude IEEE 754 32-bit floating-point value
	v1, v2, v3, v4 = float32MarshalBinaryValues(p.Message.Gps.Altitude)
	//Example altitude: float32(2.01) -> v1=0x40, v2=0x00, v3=0xa3, v4=0xd7
	p.messageOctets[9] = v1
	p.messageOctets[10] = v2
	p.messageOctets[11] = v3
	p.messageOctets[12] = v4
	// 2.0 Marshal the TimeStamp
	v1, v2, v3, v4, v5, v6, v7, v8 = p.Message.TimeStamp.marshalBytes()
	//Example timestamp: "2015-11-21T08:41:55Z"
	//	and reference time is "2305-01-01T00:00:00Z" ->
	// 		v1=0x81, v2=0x62, v3=0xf2, v4=0xa9,
	// 		v5=0x91, v6=0x2f, v7=0x7e, v8=0x00
	p.messageOctets[13] = v1
	p.messageOctets[14] = v2
	p.messageOctets[15] = v3
	p.messageOctets[16] = v4
	p.messageOctets[17] = v5
	p.messageOctets[18] = v6
	p.messageOctets[19] = v7
	p.messageOctets[20] = v8

	//hashsum ...
	// 3.0 Create HMAC: 224-bits(=28 bytes) using the
	// SHA2-256-224 hashsum algorithm ...
	//The error from createHMAC is ignored, because the code in this
	//method will make sure the error is always nil ...
	p.hMACOctets, _ = createHMAC(p.Secrets.HMACKey, p.Secrets.Salt,
		p.messageOctets)

	// 3.1 create the 'data' byte slice:
	capacity := len(p.messageOctets) + len(p.hMACOctets)
	data = make([]byte, 0, capacity)
	data = append(data, p.messageOctets...)
	data = append(data, p.hMACOctets...)

	// 4.0 Done!
	return data, nil
}

//UnmarshalBinary unmarshals a binary representation of a Payload in a byte
//slice into the data a Payload type contains.
func (p *Payload) UnmarshalBinary(data []byte) error {
	var (
		f   float32
		err error
		b   byte
		s   []byte
		//v1, v2, v3, v4 byte
	)

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// 1.0 Data argument octet size check ...
	if len(data) < PayloadOctets {
		return errors.Trace(ErrPayloadSizeTooSmall)
	}

	// 2.0 Split the 'data' slice into the message slice and the hmac slice ...
	p.messageOctets = data[:messageOctets]
	p.hMACOctets = data[messageOctets:]

	// 	3.0 HMAC checking ...
	err = checkHMAC(p.Secrets.HMACKey, p.Secrets.Salt, p.messageOctets,
		p.hMACOctets)
	if err != nil {
		annotatedErr := errors.Annotate(err,
			"HMAC checking error")
		return annotatedErr
	}

	// 4.0 Unmarshal the payload ...

	// 4.1.0 Unmarshal the GPS POI ...
	// 4.1.1 Unmarshal the FixMode ...
	b = p.messageOctets[0]
	if err = (&p.Message.Gps.FixMode).UnmarshalByte(b); err != nil {
		return errors.Annotate(err,
			"Package.Message.Gps.FixMode).UnmarshalByte error")
	}
	// 4.1.2 Unmarshal the latitude ...
	s = p.messageOctets[1:5]
	f = float32UnmarshalBinary(s)

	p.Message.Gps.Latitude = f
	// 4.1.2 Unmarshal the longitude ...
	s = p.messageOctets[5:9]
	f = float32UnmarshalBinary(s)
	p.Message.Gps.Longitude = f
	// 4.1.2 Unmarshal the altitude ...
	s = p.messageOctets[9:13]
	f = float32UnmarshalBinary(s)
	p.Message.Gps.Altitude = f

	// 4.2 Unmarshal the timestamp
	s = p.messageOctets[13:]
	p.Message.TimeStamp.unmarshalBytes(
		s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7])

	return nil
}
