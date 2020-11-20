package aes

import (
	"MiMaXue/utils"
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
	mode := cipher.NewCBCEncrypter(block,key[:block.BlockSize()])
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
	mode := cipher.NewCBCDecrypter(block,key[:block.BlockSize()])
	cipherText := make([]byte, len(data))
	mode.CryptBlocks(cipherText,data)
	cipherText = utils.ClearPKCS5Padding(cipherText,block.BlockSize())
	return cipherText,nil
}