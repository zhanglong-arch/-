package AES

import (
	"BigIntCode/utils"
	"crypto/aes"
	"crypto/cipher"
)

/**
 *
 */
func AESEncrypt(data ,key []byte,) ([]byte,error){
	block, err := aes.NewCipher(key)
	if err != nil{
		return nil,err
	}
	//明文数据尾部填充
	originData := utils.PKCS5EndPadding(data,block.BlockSize())
	//实例加密模式
	mode := cipher.NewCBCEncrypter(block,key)
	cipherText := make([]byte, len(originData))
	mode.CryptBlocks(cipherText,originData)
	return cipherText,nil
}

func AESDecrypt(data, key []byte) ([]byte,error){
	block, err := aes.NewCipher(key)
	if err != nil{
		return nil,err
	}
	//实例解密模式
	mode := cipher.NewCBCDecrypter(block,key)
	cipherText := make([]byte, len(data))
	mode.CryptBlocks(cipherText,data)
	return cipherText,nil
}