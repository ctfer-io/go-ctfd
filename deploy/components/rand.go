package components

import (
	"crypto/rand"
	"encoding/hex"
)

func randName() string {
	b := make([]byte, 4)
	rand.Read(b)
	return hex.EncodeToString(b)
}
