//Package testdata contain testdata, which are used to test package go-rmsggpsbinmsg
package testdata

//BogusData is a function that returns a byte slice, where the FixMode
//has a bogus value (0x04). The SHA hassum in BogusData is correct.
func BogusData() []byte {
	bogusMessageOctetsSlice := BogusMessageOctets()
	capacity := len(bogusMessageOctetsSlice) + len(wantBogusHmacOctetsSlice)
	s := make([]byte, 0, capacity)
	s = append(s, bogusMessageOctetsSlice...)
	s = append(s, wantBogusHmacOctetsSlice...)
	return s
}
