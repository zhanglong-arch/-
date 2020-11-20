package utils

import (
	"bytes"
)

/**
 * 为加密明文进行PKCS5尾部填充
 */

func PKCS5EndPadding(data []byte, blockSize int) []byte{

	if len(data) % blockSize == 0{
		return data
	}
	//1、计算要填充多少个
	size := blockSize - len(data) % blockSize
	//2、准备要填充的内容
	paddingText := bytes.Repeat([]byte{byte(size)},size)
	//3、填充
	return append(data,paddingText...)
}

/**
 * 去除尾部填充的数据内容，返回原数据
 * 原文： b
 * 		[98 7 7 7 7 7 7 7]
 */
func ClearPKCS5Padding(data []byte, blockSize int) []byte{
	//0、判断分组正好的情况，不需要去除尾部填充
	//if len(data)%blockSize == 0 {
	//	return data
	//}
	//1清除尾部填充
	clearSize := int(data[len(data) - 1])
	return data[:(len(data) - clearSize)]
}


/**
 * 为加密明文进行Zeros尾部填充
 */
func ZerosEndPadding(data []byte, blockSize int) []byte{

	if len(data) % blockSize == 0{
		return data
	}
	//1、计算填充多少个
	size := blockSize - len(data)%blockSize
	//2、把0填入到数据中
	paddingText := bytes.Repeat([]byte{byte(0)},size)
	return append(data,paddingText...)
}

/**
 * 将Zeros尾部填充的数据去除
 */
func ClearZerosPadding(data []byte,blockSize int) []byte{

	//if len(data) % blockSize == 0{
	//	return data
	//}
	size := blockSize - len(data) % blockSize
	return data[:len(data) - size]
}