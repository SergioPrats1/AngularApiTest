package helpers

import (
	"math/rand"
	"encoding/hex"
	"time"
)

const TokenLength = 32


func GenerateSecureToken() string {
    b := make([]byte, TokenLength)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}

func GetNewTokenExpirationDate() time.Time {
	return time.Now().Add(time.Hour * 1)
}