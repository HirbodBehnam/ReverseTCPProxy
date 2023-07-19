package util

import (
	"crypto/rand"
	"encoding/binary"
)

// RandomID will create a random ID for a connection
func RandomID() uint32 {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return binary.LittleEndian.Uint32(b)
}
