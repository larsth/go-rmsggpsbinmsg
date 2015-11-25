package binmsg

const messageOctets = timeStampOctets + gpsOctets

type Message struct {
	//TimeStamp octets: timeStampOctets(=8) bytes (type time.Duration is an int64 value)
	TimeStamp TimeStamp `json:"timestamp"`
	//Gps octets: gpsOctet bytes
	Gps Gps `json:"gps"`
}
