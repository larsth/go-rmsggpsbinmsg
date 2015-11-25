package binmsg

import (
	"encoding/binary"
	"math"
)

func float32MarshalBinary(f float32) (p []byte) {
	p = make([]byte, 4)
	u32 := math.Float32bits(f)
	binary.BigEndian.PutUint32(p, u32)
	return p
}

func float32MarshalBinaryValues(f float32) (v1, v2, v3, v4 byte) {
	p := float32MarshalBinary(f)
	return p[0], p[1], p[2], p[3]
}

func float32UnmarshalBinary(p []byte) float32 {
	u32 := binary.BigEndian.Uint32(p)
	return math.Float32frombits(u32)
}
