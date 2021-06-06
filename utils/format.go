package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func PrettyPrintStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
