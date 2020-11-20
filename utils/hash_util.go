package utils

import (
	"crypto/md5"
	"crypto/sha256"
)

/**
 * MD5哈希计算
 */
func MD5Hash(data []byte) []byte{
	md5Hash := md5.New()
	md5Hash.Write(data)
	return md5Hash.Sum(nil)
}

/**
 * SHA256哈希计算
 */
func SHA256Hash(data []byte) []byte {
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}