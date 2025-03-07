package services

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func generateToken() string {
	var seed string = string(time.Now().UnixNano())
	hash := sha256.Sum256([]byte(seed))
	return hex.EncodeToString(hash[:])
}