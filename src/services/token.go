package services

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func GenerateToken() string {
	randomNumber := rand.Int63()
	seed := strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.FormatInt(randomNumber, 10)

	hash := sha256.Sum256([]byte(seed))
	return hex.EncodeToString(hash[:])
}