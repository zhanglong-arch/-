package des

import (
	"MiMaXue/utils"
	"crypto/cipher"
	"crypto/des"
)

/**
 * 使用秘钥key对明文data进行加密
 */
func DesEncrypt (data, key []byte) ([]byte,error){
	//三要素：key、data、mode
	block, err := des.NewCipher(key)
	if err != nil{
		return nil,err
	}
	//填充后的数据
	originData := utils.PKCS5EndPadding(data,block.BlockSize())
	//实例化加密模式
	mode := cipher.NewCBCEncrypter(block,key)

	//加密
	dst := make([]byte,len(originData))
	mode.CryptBlocks(dst,originData)
	return dst,nil
}

func DesDecrypt(data, key []byte) ([]byte,error){
	block, err := des.NewCipher(key)
	if err != nil{
		return nil, err
	}
	//mode实例化
	blockMode := cipher.NewCBCDecrypter(block,key)
	//原始的，最初的：original
	originData := make([]byte, len(data))
	blockMode.CryptBlocks(originData,data)
	//对解密后的明文进行尾部填充内容去除
	originData = utils.ClearPKCS5Padding(originData,block.BlockSize())
	return originData, nil
}