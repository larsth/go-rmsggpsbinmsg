package testdata

func BogusData() []byte {
	bogusMessageOctetsSlice := BogusMessageOctets()
	capacity := len(bogusMessageOctetsSlice) + len(wantBogusHmacOctetsSlice)
	s := make([]byte, 0, capacity)
	s = append(s, bogusMessageOctetsSlice...)
	s = append(s, wantBogusHmacOctetsSlice...)
	return s
}
