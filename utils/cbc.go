package utils

import "bytes"

/**
 * 为加密明文进行PKCS5尾部填充
 */

func PKCS5EndPadding(data []byte, blockSize int) []byte{
	//1、计算要填充多少个
	size := blockSize - len(data) % blockSize
	//2、准备要填充的内容
	paddingText := bytes.Repeat([]byte{byte(size)},size)
	//3、填充
	return append(data,paddingText...)
}


/**
 * 为加密明文进行Zeros尾部填充
 */
func ZerosEndPadding(data []byte, blockSize int) []byte{
	//1、计算填充多少个
	size := blockSize - len(data)%blockSize
	//2、把0填入到数据中
	paddingText := bytes.Repeat([]byte{byte(0)},size)
	return append(data,paddingText...)
}
