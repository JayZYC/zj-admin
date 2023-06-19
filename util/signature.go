package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func ValidateSignatureWithTimestamp(accessKeyId, accessKeyIdSecret, method, version, timestamp, signature string) (bool, error) {

	mac1, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	// 对access_key进行decode
	key, err := base64.StdEncoding.DecodeString(accessKeyIdSecret)
	if err != nil {
		return false, err
	}
	var StringForSignature = fmt.Sprintf("%v\n%v\n%v\n%v", accessKeyId, timestamp, method, version)
	// 计算sign = base64(hmac_<method>(base64decode(accessKeyId), utf-8(StringForSignature)))
	h := hmac.New(sha1.New, key)
	_, err = h.Write([]byte(StringForSignature))
	if err != nil {
		return false, err
	}

	return hmac.Equal(mac1, h.Sum(nil)), nil
}
