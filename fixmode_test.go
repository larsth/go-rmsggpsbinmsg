package binmsg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

type TtdTFixModeString struct {
	Input FixMode
	Want  string
}

var tdFixModeStringWant = []*TtdTFixModeString{
	&TtdTFixModeString{
		Input: FixNotSeen,
		Want:  "Not Seen",
	},
	&TtdTFixModeString{
		Input: FixNone,
		Want:  "None",
	},
	&TtdTFixModeString{
		Input: Fix2D,
		Want:  "2D",
	},
	&TtdTFixModeString{
		Input: Fix3D,
		Want:  "3D",
	},
	&TtdTFixModeString{
		Input: FixMode(5),
		Want:  "Unknown FixMode value",
	},
}

func TestFixModeString(t *testing.T) {
	var i int = 0

	log.SetOutput(ioutil.Discard)

	for _, testItem := range tdFixModeStringWant {
		sGot := testItem.Input.String()
		if strings.Compare(testItem.Want, sGot) != 0 {
			t.Errorf("Got '%s'\nWant: '%s'\n", testItem.Want, sGot)
			t.Log("The test error ocurred at test item: tdFixModeStringWant[",
				strconv.Itoa(i), "].")
		}
		i++
	}
}

type TtdFixModeMarshal struct {
	FixMode FixMode
	WantP   []byte
	WantErr error
}

var tdFixModeMarshalJSONWant = []*TtdFixModeMarshal{
	&TtdFixModeMarshal{
		FixMode: FixNotSeen,
		WantP:   []byte{0x30},
		WantErr: nil,
	},
	&TtdFixModeMarshal{
		FixMode: FixNone,
		WantP:   []byte{0x31},
		WantErr: nil,
	},
	&TtdFixModeMarshal{
		FixMode: Fix2D,
		WantP:   []byte{0x32},
		WantErr: nil,
	},
	&TtdFixModeMarshal{
		FixMode: Fix3D,
		WantP:   []byte{0x33},
		WantErr: nil,
	},
	&TtdFixModeMarshal{
		FixMode: FixMode(byte(4)),
		WantP:   nil,
		WantErr: ErrUnknownFixMode,
	},
}

func TestFixModeMarshalJSON(t *testing.T) {
	var (
		gotP       []byte
		gotErr     error
		wantErrStr string
		gotErrStr  string
	)

	log.SetOutput(ioutil.Discard)

	for _, testItem := range tdFixModeMarshalJSONWant {
		gotP, gotErr = testItem.FixMode.MarshalJSON()
		if bytes.Compare(gotP, testItem.WantP) != 0 {
			t.Fatalf("got '%v', but want '%v'\n", gotP, testItem.WantP)
		}
		gotErrStr = fmt.Sprintf("%#v", gotErr)
		wantErrStr = fmt.Sprintf("%#v", testItem.WantErr)
		if strings.Compare(wantErrStr, gotErrStr) != 0 {
			t.Fatalf("Got '%s', but want '%s'\n", gotErrStr, wantErrStr)
		}
	}
}

type TtdFixModeUnmarshal struct {
	P           []byte
	WantFixMode FixMode
	WantErr     error
}

var tdFixModeUnmarshalJSONWant = []*TtdFixModeUnmarshal{
	&TtdFixModeUnmarshal{
		P:           []byte{0x30},
		WantFixMode: FixNotSeen,
		WantErr:     nil,
	},
	&TtdFixModeUnmarshal{
		P:           []byte{0x31},
		WantFixMode: FixNone,
		WantErr:     nil,
	},
	&TtdFixModeUnmarshal{
		P:           []byte{0x32},
		WantFixMode: Fix2D,
		WantErr:     nil,
	},
	&TtdFixModeUnmarshal{
		P:           []byte{0x33},
		WantFixMode: Fix3D,
		WantErr:     nil,
	},
	&TtdFixModeUnmarshal{
		P:           nil,
		WantFixMode: FixMode(byte(252)),
		WantErr:     ErrNilSlice,
	},
	&TtdFixModeUnmarshal{
		P:           []byte{0xf},
		WantFixMode: FixMode(byte(253)),
		WantErr:     ErrUnknownFixMode,
	},
}

func TestFixModeUnmarshalJSON(t *testing.T) {
	var (
		gotFixMode *FixMode
		gotErr     error
		wantErrStr string
		gotErrStr  string
	)

	log.SetOutput(ioutil.Discard)

	for _, testItem := range tdFixModeUnmarshalJSONWant {
		gotFixMode = new(FixMode)
		gotErr = gotFixMode.UnmarshalJSON(testItem.P)
		if byte((*gotFixMode)) != byte(testItem.WantFixMode) {
			t.Fatalf("Got: '%s', but want '%s'",
				gotFixMode.String(), testItem.WantFixMode.String())
		}
		gotErrStr = fmt.Sprintf("%#v", gotErr)
		wantErrStr = fmt.Sprintf("%#v", testItem.WantErr)
		if strings.Compare(wantErrStr, gotErrStr) != 0 {
			t.Fatalf("Got '%s', but want '%s'\n", gotErrStr, wantErrStr)
		}
	}
}

type TtdFixModeMarshalByte struct {
	FixMode  FixMode
	WantByte byte
	WantErr  error
}

var tdFixModeMarshalByteWant = []*TtdFixModeMarshalByte{
	&TtdFixModeMarshalByte{
		FixMode:  FixNotSeen,
		WantByte: byte(0),
		WantErr:  nil,
	},
	&TtdFixModeMarshalByte{
		FixMode:  FixNone,
		WantByte: byte(1),
		WantErr:  nil,
	},
	&TtdFixModeMarshalByte{
		FixMode:  Fix2D,
		WantByte: byte(2),
		WantErr:  nil,
	},
	&TtdFixModeMarshalByte{
		FixMode:  Fix3D,
		WantByte: byte(3),
		WantErr:  nil,
	},
	&TtdFixModeMarshalByte{
		FixMode:  FixMode(byte(4)),
		WantByte: byte(255),
		WantErr:  ErrUnknownFixMode,
	},
}

func TestFixModeMarshalByte(t *testing.T) {
	var (
		gotByte    byte
		gotErr     error
		wantErrStr string
		gotErrStr  string
	)

	log.SetOutput(ioutil.Discard)

	for i, testItem := range tdFixModeMarshalByteWant {
		gotByte, gotErr = testItem.FixMode.marshalByte()
		if testItem.WantByte != gotByte {
			t.Fatalf("Test %d: got '%v', but want '%v'\n",
				i, gotByte, testItem.WantByte)
		}
		gotErrStr = fmt.Sprintf("%#v", gotErr)
		wantErrStr = fmt.Sprintf("%#v", testItem.WantErr)
		if strings.Compare(wantErrStr, gotErrStr) != 0 {
			t.Fatalf("Test %d: Got '%s', but want '%s'\n", i, gotErrStr, wantErrStr)
		}
	}
}

type TtdFixModeUnmarshalByte struct {
	B           byte
	WantFixMode FixMode
	WantErr     error
}

var tdFixModeUnmarshalByteWant = []*TtdFixModeUnmarshalByte{
	//Test 0:
	&TtdFixModeUnmarshalByte{
		B:           byte(0),
		WantFixMode: FixNotSeen,
		WantErr:     nil,
	},
	//Test 1:
	&TtdFixModeUnmarshalByte{
		B:           byte(1),
		WantFixMode: FixNone,
		WantErr:     nil,
	},
	//Test 2:
	&TtdFixModeUnmarshalByte{
		B:           byte(2),
		WantFixMode: Fix2D,
		WantErr:     nil,
	},
	//Test 3:
	&TtdFixModeUnmarshalByte{
		B:           byte(3),
		WantFixMode: Fix3D,
		WantErr:     nil,
	},
	//Test 4:
	&TtdFixModeUnmarshalByte{
		B:           byte(0xf),
		WantFixMode: FixMode(byte(255)),
		WantErr:     ErrUnknownFixMode,
	},
}

func TestFixModeUnmarshalByte(t *testing.T) {
	var (
		gotFixMode *FixMode
		gotErr     error
		wantErrStr string
		gotErrStr  string
	)

	log.SetOutput(ioutil.Discard)

	for i, testItem := range tdFixModeUnmarshalByteWant {
		gotFixMode = new(FixMode)
		gotErr = gotFixMode.unmarshalByte(testItem.B)
		if byte((*gotFixMode)) != byte(testItem.WantFixMode) {
			t.Fatalf("Test: %d:: Got: '%s', but want '%s'",
				i, gotFixMode.String(), testItem.WantFixMode.String())
		}
		gotErrStr = fmt.Sprintf("%#v", gotErr)
		wantErrStr = fmt.Sprintf("%#v", testItem.WantErr)
		if strings.Compare(wantErrStr, gotErrStr) != 0 {
			t.Fatalf("Got '%s', but want '%s'\n", gotErrStr, wantErrStr)
		}
	}
}

func BenchmarkFixModeString(b *testing.B) {
	var f = Fix3D
	for i := 0; i < b.N; i++ {
		_ = f.String()
	}
}

func BenchmarkFixModeMarshalJSON(b *testing.B) {
	var f = Fix3D
	for i := 0; i < b.N; i++ {
		_, _ = f.MarshalJSON()
	}
}

func BenchmarkFixModeUnmarshalJSON(b *testing.B) {
	var (
		f = Fix3D
		p = []byte{0x33}
	)

	for i := 0; i < b.N; i++ {
		_ = f.UnmarshalJSON(p)
	}
}

func BenchmarkFixModeMarshalByte(b *testing.B) {
	var f = Fix3D
	for i := 0; i < b.N; i++ {
		_, _ = f.marshalByte()
	}
}

func BenchmarkFixModeUnmarshalByte(b *testing.B) {
	var (
		v = byte(0x03)
		f = Fix3D
	)
	for i := 0; i < b.N; i++ {
		_ = f.unmarshalByte(v)
	}
}
