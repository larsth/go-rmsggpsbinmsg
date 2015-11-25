package testdata

//Data returns a byte slice with a valid message and SHA hashsum.
func Data() []byte {
	messageOctetsSlice := MessageOctets()
	capacity := len(messageOctetsSlice) + len(wantHmacOctetsSlice)
	s := make([]byte, 0, capacity)
	s = append(s, messageOctetsSlice...)
	s = append(s, wantHmacOctetsSlice...)
	return s
}
