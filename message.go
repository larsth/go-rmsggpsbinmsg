package binmsg

const messageOctets = timeStampOctets + gpsOctets

//Message is a type that contains a TimeStamp type (when was this message
//created?), and a Gps type.
type Message struct {
	//TimeStamp octets: timeStampOctets(=8) bytes (type time.Duration is an int64 value)
	TimeStamp TimeStamp `json:"timestamp"`
	//Gps octets: gpsOctet bytes
	Gps Gps `json:"gps"`
}
