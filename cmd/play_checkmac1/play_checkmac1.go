package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"log"
	"strings"

	"github.com/larsth/go-rmsggpsbinmsg/testdata"
)

var (
	ErrNilHMACSecretKeySlice = errors.New(
		"The 'HMACSecretKey' byte slice has not been set " +
			"(It is nil or its length is zero).")
	ErrNiSaltSlice = errors.New(
		"The 'Salt' byte slice has not been set " +
			"(It is nil or its length is zero).")
	ErrNilMessageOctetsSlice = errors.New(
		"The messageoctets byte slice is nil or has a zero length.")
)

func createHMAC(hmacKey, salt, messageOctets []byte) (octets []byte, err error) {
	var (
		mac hash.Hash
	)

	octets = nil
	if len(hmacKey) == 0 {
		err = ErrNilHMACSecretKeySlice
		return
	}
	if len(salt) == 0 {
		err = ErrNiSaltSlice
		return
	}
	if messageOctets == nil {
		err = ErrNilMessageOctetsSlice
		return
	}
	err = nil

	mac = hmac.New(sha256.New224, hmacKey)

	mac.Write(hmacKey)
	mac.Write(messageOctets)
	mac.Write(salt)

	octets = mac.Sum(nil)
	return
}

func errorCheck(i int, got, want error) (s string, ok bool) {
	var (
		gotStr  string
		wantStr string
	)

	s = ""
	ok = true

	if got != nil {
		gotStr = fmt.Sprintf("%s", got)
	} else {
		gotStr = "<nil>"
	}
	if want != nil {
		wantStr = fmt.Sprintf("%s", want)
	} else {
		wantStr = "<nil>"
	}
	if strings.Compare(wantStr, gotStr) != 0 {
		s = fmt.Sprintf("Test: %d::\n Got:\n\t '%s',\n Want:\n\t '%s'\n",
			i, gotStr, wantStr)
		ok = false
	}
	return
}

func byteSliceCheck(i int, got, want []byte) (s string, ok bool) {
	if bytes.Compare(got, want) != 0 {
		s = fmt.Sprintf("Test: %d::\n Got:\n\t '%#v'\n, Want:\n\t '%#v'\n", i, got, want)
		ok = false
	} else {
		s = ""
		ok = true
	}
	return
}

func main() {
	const i int = 3
	var (
		messageOctets = []byte{
			0x3, 0x42, 0x5e, 0xc4, 0x11, 0x41, 0x49, 0xde,
			0x74, 0x40, 0x0, 0xa3, 0xd7, 0x81, 0x62, 0xf2,
			0xa9, 0x91, 0x2f, 0x7e, 0x0}
		wantHmacOctets = []byte{
			0x8e, 0x87, 0x19, 0xe9, 0xb, 0xb1, 0x44, 0x9,
			0x17, 0x56, 0x2b, 0x4b, 0x76, 0x37, 0xcd, 0x34,
			0x9c, 0x66, 0x88, 0xf4, 0x79, 0x8f, 0x8a, 0x1,
			0x1a, 0x2a, 0xa2, 0x34}
		wantErr error = nil

		//Got ...
		gotHmacOctets []byte
		gotErr        error

		s        string
		ok1, ok2 bool
	)

	gotHmacOctets, gotErr = createHMAC(testdata.HmacKey(), testdata.Salt(),
		messageOctets)
	s, ok1 = byteSliceCheck(i, gotHmacOctets, wantHmacOctets)
	if ok1 == false {
		log.Println(s)
	}
	s, ok2 = errorCheck(i, gotErr, wantErr)
	if ok2 == false {
		log.Println(s)
	}
	if ok1 == true && ok2 == true {
		fmt.Println("ok!")
	}
}
