/*
@Time : 2020/1/16 10:48
@Author : Minus4
*/
package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

var (
	Salt1 string
	Salt2 string
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func EncryptAccount(email, password string) string {
	h := md5.New()
	_, _ = io.WriteString(h, password)

	passwordMd5 := fmt.Sprintf("%x", h.Sum(nil))

	_, _ = io.WriteString(h, Salt1)
	_, _ = io.WriteString(h, email)
	_, _ = io.WriteString(h, Salt2)
	_, _ = io.WriteString(h, passwordMd5)

	return fmt.Sprintf("%x", h.Sum(nil))
}
