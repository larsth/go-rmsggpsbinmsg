package testdata

//GPS is a binary encoded GPS 3D point of interest
var gpsSlice = []byte{
	/* Fix3D */ 0x03,
	/* Latitude: float32(55.69147) -> 0x42, 0x5e, 0xc4, 0x11 */
	0x42, 0x5e, 0xc4, 0x11,
	/* Longitude: float32(12.61681) -> 0x41, 0x49, 0xde, 0x74 */
	0x41, 0x49, 0xde, 0x74,
	/* Altitude: float32(2.01) -> 0x40, 0x00, 0xa3, 0xd7 */
	0x40, 0x00, 0xa3, 0xd7,
}

//GPS retuns a byte slice with the GPS POI:
//FixMode: Fix3D, Latitude: 55.69147, Longitude: 12.61681, Altitude: 2.01
func GPS() []byte {
	s := make([]byte, 0, len(gpsSlice))
	s = append(s, gpsSlice...)
	return s
}

//ts is a binary encoded time, where the refence time is
//"2305-01-01T00:00:00Z" (RFC3339 encoded string).
//A dataTimeStamp is a 64-bit 2's complement big-endian encoded
//value of nanoseconds relative to the reference time:
var ts = []byte{
	/* "2015-11-21T08:41:55Z" -> 0x81, 0x62, 0xf2, 0xa9, 0x91, 0x2f, 0x7e, 0x00 */
	0x81, 0x62, 0xf2, 0xa9, 0x91, 0x2f, 0x7e, 0x00,
}

//TimeStamp returns a byte slice with a binary encoded time, which is relative
//to the refence time which is "2305-01-01T00:00:00Z" (RFC3339 encoded string).
//It is a 64-bit 2's complement big-endian encoded value of nanoseconds
//relative to the reference time.
func TimeStamp() []byte {
	s := make([]byte, 0, len(ts))
	s = append(s, ts...)
	return s
}

//MessageOctets returns a byte slice, which contains GPS() immediately followed
//by TimeStamp().
func MessageOctets() []byte {
	s := make([]byte, 0, 256)
	s = append(s, gpsSlice...)
	timeStampSlice := TimeStamp()
	s = append(s, timeStampSlice...)
	return s
}

var bogusGpsSlice = []byte{
	/*FixMode 3D: changed to bogus value 0x04 */
	0x04,
	/* Latitude: float32(55.69147) -> 0x42, 0x5e, 0xc4, 0x11, but 1st byte
	changed to bugus value: 0xaa */
	0x42, 0x5e, 0xc4, 0x11,
	/* Longitude: float32(12.61681) -> 0x41, 0x49, 0xde, 0x74 */
	0x41, 0x49, 0xde, 0x74,
	/* Altitude: float32(2.01) -> 0x40, 0x00, 0xa3, 0xd7 */
	0x40, 0x00, 0xa3, 0xd7,
}

//BogusMessageOctets returns a byte slice, which contains BogusGPS() immediately
// followed by TimeStamp().
func BogusMessageOctets() []byte {
	s := make([]byte, 0, 256)
	s = append(s, bogusGpsSlice...)
	timeStampSlice := TimeStamp()
	s = append(s, timeStampSlice...)
	return s
}
