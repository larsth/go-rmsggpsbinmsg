package binmsg

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/larsth/go-gpsfix"
	"github.com/larsth/go-rmsggpsbinmsg/testdata"
)

type TtdCreteHmac struct {
	//Input ...
	HmacKey       []byte
	Salt          []byte
	MessageOctets []byte

	//Want ...
	WantHmacOctets []byte
	WantErr        error
}

var (
	tdCreateHmac = []*TtdCreteHmac{
		//Test 0:
		&TtdCreteHmac{
			HmacKey:        nil,
			Salt:           nil,
			MessageOctets:  nil,
			WantHmacOctets: nil,
			WantErr:        ErrNilHMACSecretKeySlice,
		},
		//Test 1:
		&TtdCreteHmac{
			HmacKey:        testdata.HmacKey(),
			Salt:           nil,
			MessageOctets:  nil,
			WantHmacOctets: nil,
			WantErr:        ErrNiSaltSlice,
		},
		//Test 2:
		&TtdCreteHmac{
			HmacKey:        testdata.HmacKey(),
			Salt:           testdata.Salt(),
			MessageOctets:  nil,
			WantHmacOctets: nil,
			WantErr:        ErrNilMessageOctetsSlice,
		},
	}
)

func TestCreateHMAC1(t *testing.T) {
	var (
		//Got ...
		gotHmacOctets []byte
		gotErr        error

		s  string
		ok bool
	)
	for i, testItem := range tdCreateHmac {
		gotHmacOctets, gotErr = createHMAC(testItem.HmacKey, testItem.Salt,
			testItem.MessageOctets)
		s, ok = byteSliceCheck(i, gotHmacOctets, testItem.WantHmacOctets)
		if ok == false {
			t.Error(s)
		}
		s, ok = errorCheck(i, gotErr, testItem.WantErr)
		if ok == false {
			t.Error(s)
		}
	}
}

func TestCreateHMAC2(t *testing.T) {
	const i int = 3
	var (
		wantErr error // = nil

		//Got ...
		gotHmacOctets []byte
		gotErr        error

		s  string
		ok bool
	)

	gotHmacOctets, gotErr = createHMAC(testdata.HmacKey(), testdata.Salt(),
		testdata.MessageOctets())
	s, ok = byteSliceCheck(i, gotHmacOctets, testdata.WantHmacOctets())
	if ok == false {
		t.Error(s)
	}
	s, ok = errorCheck(i, gotErr, wantErr)
	if ok == false {
		t.Error(s)
	}
}

type TtdCheckMac struct {
	//Input:
	HmacOctets    []byte
	HmacKey       []byte
	Salt          []byte
	MessageOctets []byte

	//Want:
	WantHmacOctets []byte
	WantErr        error
}

var tdCheckHMAC = []*TtdCheckMac{
	//Test 0:
	&TtdCheckMac{
		HmacOctets:     nil,
		HmacKey:        nil,
		Salt:           nil,
		MessageOctets:  nil,
		WantHmacOctets: nil,
		WantErr:        ErrHMACOctetsWrongSize,
	},
	//Test 1:
	&TtdCheckMac{
		HmacOctets:     testdata.WantHmacOctets(),
		HmacKey:        nil,
		Salt:           nil,
		MessageOctets:  nil,
		WantHmacOctets: nil,
		//WantErr: ErrNilHMACSecretKeySlice inlined via juju/errors package
		WantErr: fmt.Errorf("%s%s",
			"expectedMAC: createHMAC error: The 'HMACSecretKey' byte slice ",
			"has not been set (It is nil or its length is zero)."),
	},
	//Test 2:
	&TtdCheckMac{
		HmacOctets:     testdata.WantHmacOctets(),
		HmacKey:        testdata.HmacKey(),
		Salt:           nil,
		MessageOctets:  nil,
		WantHmacOctets: nil,
		//WantErr: ErrNiSaltSlice inlined via juju/errors package
		WantErr: fmt.Errorf("%s%s",
			"expectedMAC: createHMAC error: The 'Salt' byte slice has ",
			"not been set (It is nil or its length is zero)."),
	},
	//Test 3:
	&TtdCheckMac{
		HmacOctets:     testdata.WantHmacOctets(),
		HmacKey:        testdata.HmacKey(),
		Salt:           testdata.Salt(),
		MessageOctets:  nil,
		WantHmacOctets: nil,
		//WantErr: ErrNilMessageOctetsSlice inlined via juju/errors package
		WantErr: fmt.Errorf("%s%s",
			"expectedMAC: createHMAC error: The messageoctets byte slice ",
			"is nil or has a zero length."),
		//,
	},
	//Test 4, check should fail (FixMode changed from 0x04 to 0x05):
	&TtdCheckMac{
		HmacOctets:     testdata.WantHmacOctets(),
		HmacKey:        testdata.HmacKey(),
		Salt:           testdata.Salt(),
		MessageOctets:  testdata.BogusMessageOctets(),
		WantHmacOctets: nil,
		WantErr:        ErrHMACcheckFailed,
	},
	//Test 5, check should be successful (no error):
	&TtdCheckMac{
		HmacOctets:     testdata.WantHmacOctets(),
		HmacKey:        testdata.HmacKey(),
		Salt:           testdata.Salt(),
		MessageOctets:  testdata.MessageOctets(),
		WantHmacOctets: testdata.WantHmacOctets(),
		WantErr:        nil,
	},
}

func TestCheckHMAC(t *testing.T) {
	var (
		gotErr error
		s      string
		ok     bool
	)
	for i, testItem := range tdCheckHMAC {
		gotErr = checkHMAC(testItem.HmacKey, testItem.Salt,
			testItem.MessageOctets, testItem.HmacOctets)
		s, ok = errorCheck(i, gotErr, testItem.WantErr)
		if ok == false {
			t.Error(s)
		}
	}
}

type TWantFunc func(i int, testItem *TtdPayloadInit, p *Payload) (s string, ok bool)

type TtdPayloadInit struct {
	//Input:
	HmacKey []byte
	Salt    []byte

	//Want:
	WantErr  error
	WantFunc TWantFunc
}

func funcPayloadInitWant(i int, testItem *TtdPayloadInit, p *Payload) (s string, ok bool) {
	s, ok = byteSliceCheck(i, p.Secrets.HMACKey, testItem.HmacKey)
	if ok == false {
		return
	}
	s, ok = byteSliceCheck(i, p.Secrets.Salt, testItem.Salt)
	if ok == false {
		return
	}
	s = ""
	ok = true
	return
}

var tdPayloadInit1 = []*TtdPayloadInit{
	//Test 0:
	&TtdPayloadInit{
		HmacKey: nil,
		Salt:    nil,
		//WantErr: ErrNilHMACSecretKeySlice inlined via juju/errors package
		WantErr: fmt.Errorf("%s",
			"The 'HMACSecretKey' byte slice has not been set (It is nil or its length is zero)."),
		WantFunc: nil,
	},
	//Test 1:
	&TtdPayloadInit{
		HmacKey: testdata.HmacKey(),
		Salt:    nil,
		//WantErr: ErrNiSaltSlice inlined via juju/errors package
		WantErr: fmt.Errorf("%s",
			"The 'Salt' byte slice has not been set (It is nil or its length is zero)."),
		WantFunc: nil,
	},
	//Test 2:
	&TtdPayloadInit{
		HmacKey:  testdata.HmacKey(),
		Salt:     testdata.Salt(),
		WantErr:  nil,
		WantFunc: funcPayloadInitWant,
	},
}

func TestPayloadInit(t *testing.T) {
	var (
		p      *Payload
		gotErr error
		s      string
		ok     bool
	)
	for i, testItem := range tdPayloadInit1 {
		p = new(Payload)
		gotErr = p.Init(testItem.HmacKey, testItem.Salt)
		if s, ok = errorCheck(i, gotErr, testItem.WantErr); ok == false {
			t.Error(s)
		}
		if testItem.WantFunc != nil {
			if s, ok = testItem.WantFunc(i, testItem, p); ok == false {
				t.Error(s)
			}
		}
	}
}

var tdPayloadInit2 = []*TtdPayloadInit{
	//Test 0:
	&TtdPayloadInit{
		HmacKey: nil,
		Salt:    nil,
		//WantErr: ErrNilHMACSecretKeySlice inlined via juju/errors package
		WantErr: fmt.Errorf("%s%s",
			"Init(hmacKey, salt) error: The 'HMACSecretKey' ",
			"byte slice has not been set (It is nil or its length is zero)."),
		WantFunc: nil,
	},
	//Test 1:
	&TtdPayloadInit{
		HmacKey: testdata.HmacKey(),
		Salt:    nil,
		//WantErr: ErrNiSaltSlice inlined via juju/errors package
		WantErr: fmt.Errorf("%s%s",
			"Init(hmacKey, salt) error: The 'Salt' byte slice has not been set",
			" (It is nil or its length is zero)."),
		WantFunc: nil,
	},
	//Test 2:
	&TtdPayloadInit{
		HmacKey:  testdata.HmacKey(),
		Salt:     testdata.Salt(),
		WantErr:  nil,
		WantFunc: funcPayloadInitWant,
	},
}

func TestPayloadNew(t *testing.T) {
	var (
		p      *Payload
		gotErr error
		s      string
		ok     bool
	)
	for i, testItem := range tdPayloadInit2 {
		p, gotErr = New(testItem.HmacKey, testItem.Salt)
		if s, ok = errorCheck(i, gotErr, testItem.WantErr); ok == false {
			t.Error(s)
		}
		if testItem.WantFunc != nil {
			if s, ok = testItem.WantFunc(i, testItem, p); ok == false {
				t.Error(s)
			}
		}
	}
}

type TtdPayloadMarshalBinary struct {
	//Init
	Init func(testItem *TtdPayloadMarshalBinary)
	P    *Payload

	//Want:
	WantData []byte
	WantErr  error
}

var tdPayloadMarshalBinary = []*TtdPayloadMarshalBinary{
	//Test 0:
	&TtdPayloadMarshalBinary{
		Init: func(testItem *TtdPayloadMarshalBinary) {
			testItem.P = tFnCreatePayload()
			testItem.P.Secrets.HMACKey = nil
		},
		WantData: nil,
		WantErr:  ErrNilHMACSecretKeySlice,
	},
	//Test 1:
	&TtdPayloadMarshalBinary{
		Init: func(testItem *TtdPayloadMarshalBinary) {
			testItem.P = tFnCreatePayload()
			testItem.P.Secrets.Salt = nil
		},
		WantData: nil,
		WantErr:  ErrNiSaltSlice,
	},
	//Test 2:
	&TtdPayloadMarshalBinary{
		Init: func(testItem *TtdPayloadMarshalBinary) {
			testItem.P = tFnCreatePayload()
			fmPtr := &testItem.P.Message.Gps.FixMode
			(*fmPtr) = gpsfix.FixMode(0x04)
		},
		WantData: nil,
		//WantErr: gpsfix.ErrUnknownFixMode inlined via juju/errors package
		WantErr: fmt.Errorf("%s%s%s",
			"Payload.Message.Gps.FixMode.MarshalByte() error: Unknown FixMode ",
			"value: valid values are FixNotSeen, 0 or FixNone, 1 or Fix2D, 2 ",
			"or Fix3D, 3."),
	},
	//Test 3:
	&TtdPayloadMarshalBinary{
		Init: func(testItem *TtdPayloadMarshalBinary) {
			testItem.P = tFnCreatePayload()
			testItem.WantData = testdata.Data()
		},
		WantErr: nil,
	},
}

func tFnCreatePayload() (p *Payload) {
	p, _ = New(testdata.HmacKey(), testdata.Salt())
	fmPtr := &(p.Message.Gps.FixMode)
	(*fmPtr) = gpsfix.Fix3D
	p.Message.Gps.Latitude = testdata.Latitude()
	p.Message.Gps.Longitude = testdata.Longitude()
	p.Message.Gps.Altitude = testdata.Altitude()
	p.Message.TimeStamp.Time = testdata.Time()
	return p
}

func TestPayloadMarshalBinary(t *testing.T) {
	var (
		gotData []byte
		gotErr  error
		s       string
		ok      bool
	)
	for i, testItem := range tdPayloadMarshalBinary {
		testItem.Init(testItem)
		gotData, gotErr = testItem.P.MarshalBinary()
		if s, ok = byteSliceCheck(i, gotData, testItem.WantData); ok == false {
			t.Error(s)
		}
		if s, ok = errorCheck(i, gotErr, testItem.WantErr); ok == false {
			t.Error(s)
		}
	}
}

func TestPayloadUnmarshalBinary0(t *testing.T) {
	var (
		wantErr = ErrPayloadSizeTooSmall
		data    []byte
		p       = new(Payload)
		i       int // = 0
		gotErr  error
		s       string
		ok      bool
	)

	//Init ...
	p.Init(testdata.HmacKey(), testdata.Salt())
	data = testdata.MessageOctets()

	gotErr = p.UnmarshalBinary(data)
	if s, ok = errorCheck(i, gotErr, wantErr); ok == false {
		t.Error(s)
	}
}

func TestPayloadUnmarshalBinary1(t *testing.T) {
	var (
		//wantErr: ErrHMACcheckFailed inlined via juju/errors package
		wantErr = fmt.Errorf("%s",
			"HMAC checking error: SECURITY: HMAC check failed")
		data   []byte
		p      = new(Payload)
		i      int // = 0
		gotErr error
		s      string
		ok     bool
	)

	//Init ...
	p.Init(testdata.HmacKey(), testdata.Salt())
	data = testdata.Data()

	//Change the byte at index 1 in byte slice 'data' to a bogus value ...
	data[1] = 0xaa

	gotErr = p.UnmarshalBinary(data)
	if s, ok = errorCheck(i, gotErr, wantErr); ok == false {
		t.Error(s)
	}
}

func TestPayloadUnmarshalBinary2(t *testing.T) {
	var (
		//wantErr: gpsfix.ErrUnknownFixMode inlined via juju/errors package
		wantErr = fmt.Errorf("%s%s%s",
			"Package.Message.Gps.FixMode).UnmarshalByte error: ",
			"Unknown FixMode value: valid values are FixNotSeen, 0 or ",
			"FixNone, 1 or Fix2D, 2 or Fix3D, 3.")
		data   []byte
		p      = new(Payload)
		i      int // = 0
		gotErr error
		s      string
		ok     bool
	)

	//Init ...
	p.Init(testdata.HmacKey(), testdata.Salt())
	data = testdata.BogusData()

	gotErr = p.UnmarshalBinary(data)
	if s, ok = errorCheck(i, gotErr, wantErr); ok == false {
		t.Error(s)
	}
}

func TestPayloadUnmarshalBinary3(t *testing.T) {
	var (
		wantErr        error // = nil
		data           []byte
		p              = new(Payload)
		i              int // = 0
		gotErr         error
		s, sGot, sWant string
		ok             bool
		tGot, tWant    time.Time
	)

	//Init ...
	p.Init(testdata.HmacKey(), testdata.Salt())
	data = testdata.Data()

	gotErr = p.UnmarshalBinary(data)
	if s, ok = errorCheck(i, gotErr, wantErr); ok == false {
		t.Error(s)
	}
	if ok == true {
		sGot = p.Message.Gps.FixMode.String()
		fmPtr := new(gpsfix.FixMode)
		(*fmPtr) = gpsfix.Fix3D
		sWant = fmPtr.String()
		if strings.Compare(sWant, sGot) != 0 {
			s = mkStrErrString("FixMode", sGot, sWant)
			t.Error(s)
		}

		if p.Message.Gps.Latitude != testdata.Latitude() {
			sGot = strconv.FormatFloat(p.Message.Gps.Lat(), 'f', -1, 32)
			sWant = strconv.FormatFloat(float64(testdata.Latitude()), 'f', -1, 64)
			s = mkStrErrString("GPS, Latitude", sGot, sWant)
			t.Error(s)
		}

		if p.Message.Gps.Longitude != testdata.Longitude() {
			sGot = strconv.FormatFloat(p.Message.Gps.Lon(), 'f', -1, 32)
			sWant = strconv.FormatFloat(float64(testdata.Longitude()), 'f', -1, 64)
			s = mkStrErrString("GPS, Longitude", sGot, sWant)
			t.Error(s)
		}

		if p.Message.Gps.Altitude != testdata.Altitude() {
			sGot = strconv.FormatFloat(p.Message.Gps.Alt(), 'f', -1, 32)
			sWant = strconv.FormatFloat(float64(testdata.Altitude()), 'f', -1, 64)
			s = mkStrErrString("GPS, Altitude", sGot, sWant)
			t.Error(s)
		}

		tGot = p.Message.TimeStamp.Time
		tWant = testdata.Time()
		if tWant.Equal(tGot) == false {
			s = mkStrErrString("'TimeStamp time'",
				tGot.String(), tWant.String())
			t.Error(s)
		}
	}
}

func BenchmarkPayloadCreateHMAC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = createHMAC(testdata.HmacKey(), testdata.Salt(),
			testdata.MessageOctets())
	}
}

func BenchmarkPayloadCheckHMAC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = checkHMAC(testdata.HmacKey(), testdata.Salt(),
			testdata.MessageOctets(), testdata.WantHmacOctets())
	}
}

func BenchmarkPayloadInit(b *testing.B) {
	var p = new(Payload)
	for i := 0; i < b.N; i++ {
		_ = p.Init(testdata.HmacKey(), testdata.Salt())
	}
}

func BenchmarkPayloadNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = New(testdata.HmacKey(), testdata.Salt())
	}
}

func BenchmarkPayloadMarshalBinary(b *testing.B) {
	var p = tFnCreatePayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = p.MarshalBinary()
	}
}

func BenchmarkPayloadUnmarshalBinary(b *testing.B) {
	var p = new(Payload)
	var data = testdata.Data()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.UnmarshalBinary(data)
	}
}
