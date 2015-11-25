package binmsg

import (
	"strconv"
	"testing"
)

func TestFloat32MarshalBinary(t *testing.T) {
	want := []byte{0x42, 0x5e, 0xc4, 0x11}
	got := float32MarshalBinary(float32(55.69147))
	if s, ok := byteSliceCheck(0, got, want); ok == false {
		t.Fatal(s)
	}
}

func TestFloat32MarshalBinaryValues(t *testing.T) {
	want := []byte{0x42, 0x5e, 0xc4, 0x11}
	gotV1, gotV2, gotV3, gotV4 := float32MarshalBinaryValues(float32(55.69147))
	got := make([]byte, 4)
	got[0] = gotV1
	got[1] = gotV2
	got[2] = gotV3
	got[3] = gotV4
	if s, ok := byteSliceCheck(0, got, want); ok == false {
		t.Fatal(s)
	}
}

func TestFloat32UnmarshalBinary(t *testing.T) {
	p := []byte{0x42, 0x5e, 0xc4, 0x11}
	want := float32(55.69147)
	got := float32UnmarshalBinary(p)
	if want != got {
		gotStr := strconv.FormatFloat(float64(got), 'f', -1, 32)
		wantStr := strconv.FormatFloat(float64(got), 'f', -1, 32)
		s := mkStrErrString("", gotStr, wantStr)
		t.Fatal(s)
	}
}

func BenchmarkMiscFloat32MarshalBinary(b *testing.B) {
	var f = float32(123.45)
	for i := 0; i < b.N; i++ {
		_ = float32MarshalBinary(f)
	}
}

func BenchmarkMiscFloat32MarshalBinaryValues(b *testing.B) {
	var f = float32(123.45)
	for i := 0; i < b.N; i++ {
		_, _, _, _ = float32MarshalBinaryValues(f)
	}
}

func BenchmarkMiscFloat32UnmarshalBinary(b *testing.B) {
	var p = []byte{0x01, 0x23, 0x45, 0x67}
	for i := 0; i < b.N; i++ {
		_ = float32UnmarshalBinary(p)
	}
}
