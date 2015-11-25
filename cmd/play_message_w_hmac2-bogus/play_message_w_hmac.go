package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"strings"

	"github.com/larsth/go-rmsggpsbinmsg/testdata"
)

func main() {
	var (
		data       []byte
		message    []byte
		mac        hash.Hash
		hmacOctets []byte
	)

	//message with timestamp and hmac zlib writer test program

	fmt.Printf("dataGPS: %#v\n", testdata.GPS)
	fmt.Println(strings.Repeat("-", 80))

	fmt.Printf("dataTimeStamp: %#v\n", testdata.TimeStamp)
	fmt.Println(strings.Repeat("-", 80))

	message = testdata.BogusMessageOctets()

	fmt.Printf("message: %#v\n", message)
	fmt.Println(strings.Repeat("-", 80))

	data = make([]byte, 0, 256)
	data = append(data, message...)

	//HMAC ......
	mac = hmac.New(sha256.New224, testdata.HmacKey())
	mac.Write(testdata.HmacKey())
	mac.Write(message)
	mac.Write(testdata.Salt())
	hmacOctets = mac.Sum(nil)
	data = append(data, hmacOctets...)
	//end HMAC

	fmt.Printf("hmac octets: %#v\n", hmacOctets)
	fmt.Println(strings.Repeat("-", 80))

	fmt.Printf("message+HMAC byte slice: %#v\n", data)
	fmt.Printf("# of bytes in the message+HMAC byte slice: %d\n", len(data))
	fmt.Println(strings.Repeat("-", 80))
}
