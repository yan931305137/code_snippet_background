package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(password string) (string, error) {
	// 第一次MD5加密
	password, err := Encrypt(password)
	if err != nil {
		return "", err
	}
	// 第二次MD5加密
	password2, err := Encrypt(password)
	if err != nil {
		return "", err
	}
	return password2, nil
}

func Encrypt(data interface{}) (encrypt string, err error) {
	if data == nil {
		return "", nil
	}
	switch value := data.(type) {
	case string:
		return EncryptBytes([]byte(value))
	case []byte:
		return EncryptBytes(value)
	default:
		return "", nil
	}
}

func EncryptBytes(data []byte) (encrypt string, err error) {
	h := md5.New()
	if _, err = h.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
