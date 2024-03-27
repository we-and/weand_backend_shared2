package business
import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func GenerateEmailConfirmToken(email string) string {
	res := email + time.Now().Format("2006-01-02 15:04:05")
	data := []byte(res)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
