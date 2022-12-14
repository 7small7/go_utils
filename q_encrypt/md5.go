package q_encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

func emptyStrVerify(encryptStr string) (string, error) {
	if encryptStr == "" {
		return "", errors.New("加密字符串不能为空")
	}
	return encryptStr, nil
}

// Md5Encrypt 获取小写32位加密字符串
func Md5Encrypt(encryptStr string) (string, error) {
	_, err := emptyStrVerify(encryptStr)
	if err != nil {
		return "", err
	}
	str := md5.Sum([]byte(encryptStr))
	md5str := fmt.Sprintf("%x", str)
	return md5str, nil
}

// Md516Encrypt 获取小写16位加密字符串
func Md516Encrypt(encryptStr string) (string, error) {
	_, err := emptyStrVerify(encryptStr)
	if err != nil {
		return "", err
	}
	h := md5.New()
	h.Write([]byte(encryptStr))
	return hex.EncodeToString(h.Sum(nil))[0:16], nil
}

// Md5EncryptUpper 获取大写32位加密字符串
func Md5EncryptUpper(encryptStr string) (string, error) {
	_, err := emptyStrVerify(encryptStr)
	if err != nil {
		return "", err
	}
	encrypt, err := Md5Encrypt(encryptStr)
	return strings.ToUpper(encrypt), err
}

// Md516EncryptUpper 获取大写16位加密字符串
func Md516EncryptUpper(encryptStr string) (string, error) {
	_, err := emptyStrVerify(encryptStr)
	if err != nil {
		return "", err
	}
	encrypt, err := Md516Encrypt(encryptStr)
	return strings.ToUpper(encrypt), err
}
