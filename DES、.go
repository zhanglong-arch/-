package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {

	key := []byte("c1906031")
	data := "强者自救，圣者渡人，弱者自灭"
	block, err := des.NewCipher(key)
	if err !=nil{
		panic(err.Error())
	}

	//计算需要填充多少
	paddingSize := block.BlockSize() - len([]byte(data))%block.BlockSize()
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	dataText := append([]byte(data), paddingText...)
		mode := cipher.NewCBCEncrypter(block,key)
		dst := make([]byte,len([]byte(dataText)))
		mode.CryptBlocks(dst, []byte(dataText))
		fmt.Println("加密后的内容：",string(dst))

		//二、接收端解密
		//DES解密：
		key1 := []byte("c1906031")

		blcok1, err := des.NewCipher(key1)
		if err != nil{
			panic(err.Error())
		}

		//密文数据
		cipherData := dst

		//实例化一个解密模式实例
		blockMode1 := cipher.NewCBCDecrypter(blcok1,key1)

		//创建明文容器
		originalData := make([]byte, len(cipherData))
		//解密
		blockMode1.CryptBlocks(originalData,cipherData)
		fmt.Println("解密后的内容：",)
}