package main

import (
	"fmt"

	binmsg "github.com/larsth/go-rmsggpsbinmsg"
	"github.com/larsth/go-rmsggpsbinmsg/testdata"
)

func main() {
	const (
		timeStampOctets = 8
		gpsOctets       = 13
		messageOctets   = timeStampOctets + gpsOctets
	)
	var (
		data []byte
		p    *binmsg.Payload
	)
	data = testdata.Data()

	p = new(binmsg.Payload)

	p.MessageOctets = data[:messageOctets]
	fmt.Printf("%#v\n", p.MessageOctets)
	p.HMACOctets = data[messageOctets:]
	fmt.Printf("%#v\n", p.HMACOctets)
}
