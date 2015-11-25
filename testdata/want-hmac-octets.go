package testdata

var wantHmacOctetsSlice = []byte{
	0x1b, 0xb6, 0xc1, 0x9f, 0xfe, 0x22, 0xc2, 0x70,
	0x4b, 0xfd, 0x44, 0x6c, 0x5b, 0x3a, 0x2f, 0x7b,
	0xa6, 0x9b, 0x47, 0x5a, 0x1a, 0x84, 0x87, 0x4b,
	0xb0, 0x58, 0xb, 0x29}

var wantBogusHmacOctetsSlice = []byte{
	0x9c, 0xe7, 0xa5, 0xc, 0x60, 0xb1, 0xc7, 0x65,
	0xef, 0x9b, 0xaf, 0xa0, 0xec, 0xa7, 0xcb, 0x20,
	0x53, 0xe, 0x5a, 0xfb, 0x80, 0x53, 0x8, 0x5a,
	0xc1, 0x35, 0xd, 0xc3}

//WantHmacOctets returns a byte slice - used in TestCheckHMAC (test #5),
//TestCreateHMAC2, and BenchmarkPayloadCheckHMAC.
func WantHmacOctets() []byte {
	s := make([]byte, 0, len(wantHmacOctetsSlice))
	s = append(s, wantHmacOctetsSlice...)
	return s
}

//WantBogusHmacOctets returns a byte slice
func WantBogusHmacOctets() []byte {
	s := make([]byte, 0, len(wantBogusHmacOctetsSlice))
	s = append(s, wantBogusHmacOctetsSlice...)
	return s
}
