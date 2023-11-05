package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	md5Hash := hex.EncodeToString(hash[:])
	return md5Hash
}
