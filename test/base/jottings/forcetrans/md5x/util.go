package md5x

import (
	"crypto/md5"
	"encoding/hex"
)

func HexStr(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}
