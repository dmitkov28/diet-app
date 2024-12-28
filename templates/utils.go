package templates

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashStr(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashedBytes := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}
