package binmsg

import "errors"

var (
	ErrUnknownFixMode = errors.New("Unknown FixMode value: valid values " +
		"are FixNotSeen, 0 or FixNone, 1 or Fix2D, 2 or Fix3D, 3.")
	ErrNilSlice = errors.New("Nil input slice to 'func (f *FixMode) " +
		"UnmarshalJSON(p []byte) error'.")

	//Payload

	ErrHMACOctetsWrongSize = errors.New(
		"The Payload'sHMACOctets slice does not have a slice of " +
			"sha256.Size224 octets in length.")
	ErrHMACcheckFailed = errors.New(
		"SECURITY: HMAC check failed.")
	ErrPayloadSizeTooSmall   error
	ErrNilHMACSecretKeySlice = errors.New(
		"The 'HMACSecretKey' byte slice has not been set " +
			"(It is nil or its length is zero).")
	ErrNiSaltSlice = errors.New(
		"The 'Salt' byte slice has not been set " +
			"(It is nil or its length is zero).")
	ErrNilMessageOctetsSlice = errors.New(
		"The messageoctets byte slice is nil or has a zero length.")
	ErrNilSSlice = errors.New(
		"The 's' byte slice is nil.")
)
