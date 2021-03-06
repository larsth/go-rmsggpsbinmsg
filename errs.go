//Package binmsg can marshal and unmarshal a binary encoded message with
//a timestamp (time, data) and GPS data: Lat., Lon., Alt., FixMode, and a
//HMAC hashsum.
package binmsg

import "errors"

var (
	//Payload

	//ErrHMACOctetsWrongSize is the error for HMACOctet with a wrong size.
	ErrHMACOctetsWrongSize = errors.New(
		"The Payload'sHMACOctets slice does not have a slice of " +
			"sha256.Size224 octets in length.")

	//ErrHMACcheckFailed is the error for failed SECURITY check.
	//The payload does _NOT_ has vaild data.
	ErrHMACcheckFailed = errors.New(
		"SECURITY: HMAC check failed")

	//ErrPayloadSizeTooSmall is the error for a payload byte slice with an
	//incorrect length.
	//(The initBinMsg function creates this error).
	ErrPayloadSizeTooSmall error

	//ErrNilHMACSecretKeySlice is the error for a nil or zero length HMACkey secret.
	ErrNilHMACSecretKeySlice = errors.New(
		"The 'HMACSecretKey' byte slice has not been set " +
			"(It is nil or its length is zero).")

	//ErrNiSaltSlice is the error for a nil or zero length salt secret.
	ErrNiSaltSlice = errors.New(
		"The 'Salt' byte slice has not been set " +
			"(It is nil or its length is zero).")

	//ErrNilMessageOctetsSlice is the error for a message byte slice with a
	//wrong length/size.
	ErrNilMessageOctetsSlice = errors.New(
		"The messageoctets byte slice is nil or has a zero length.")
	ErrNilByteSlice = errors.New(
		"<nil> byte slice")
)
