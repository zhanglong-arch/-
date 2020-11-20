package main

import (
	_des "MiMaXue/3des"
	"MiMaXue/aes"
	"MiMaXue/des"
	"MiMaXue/ecc"
	"fmt"
)

func main(){
	//privateKey,err := rsa.CreateRSAPairKeys()
	//
	//data6 := []byte("你是好人啊!")
	//fmt.Println(rsa.RSASign(privateKey,data6))
	//return

	//一、des算法：key、data
	key := []byte("c1906031")
	data := "都走了"
	fmt.Println("加密前：",data)
	//1、加密
	enrs, _:= des.DesEncrypt([]byte(data),key)
	fmt.Println("加密后：",string(enrs))
	//2、解密
	ders, _ := des.DesDecrypt(enrs,key)
	fmt.Println("解密后：",string(ders))

	//二、3DES算法
	key1 := []byte("202011122020111220201112")//3des密钥长度必须为24字节
	data1 := "窗含西岭千秋雪，门泊东吴万里船"

	cipherText1, err := _des.TripleDesEncrypt([]byte(data1),key1)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("3des加密后的内容：",string(cipherText1))
	originalText1, err := _des.TripleDesDecrypt(cipherText1,key1)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("3des解密密后的内容：",string(originalText1))

	//3、AES算法
	//AES算法密钥长度必须是：16字节、24字节、32字节
	//16->128位	24->192位	32->256位
	key2 := []byte("20201112202011122020111220201112")//8
	data2 := "只因为在人群中多看了你一眼，再也没能忘记你容颜"
	cipherText2, err := aes.AESEncrypt([]byte(data2),key2)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("AES算法加密后的内容：",string(cipherText2))
	originalText2, err := aes.AESDecrypt(cipherText2,key2)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("AES算法解密后的内容：",string(originalText2))

	//4、RSA算法
	fmt.Println("================RSA算法==================")
	//data3 := "在天愿做比翼鸟，大难临头各自飞"
	////4.1 生成一对密钥
	//pri, err := rsa.CreateRSAPairKeys()
	//if err != nil {
	//	fmt.Println("rsa算法密钥生成失败：",err.Error())
	//	return
	//}

	//4.1.5将私钥保存到文件中
	//err = rsa.GeneratePriPem(pri,"gg")
	//if err != nil{
	//	fmt.Println("私钥证书文件生成失败")
	//	return
	//}
	//err = rsa.GeneratePubPem(pri.PublicKey,"pzz")
	//if err != nil {
	//	fmt.Println("公钥证书文件生成失败")
	//}
	//err = rsa.GenerateKeysPem("xw")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}



	//4.2 使用生成的密钥对数据进行加密
	//cipherText3, err := rsa.RSAEncrypt(pri.PublicKey, []byte(data3))
	//if err != nil{
	//	fmt.Println("rsa算法加密失败：",err.Error())
	//}
	//fmt.Println("rsa算法加密成功：",string(cipherText3))
	//
	////4.3 使用私钥进行解密
	//originalText3, err := rsa.RSADecrypt(pri, cipherText3)
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println("rsa算法解密成功：",string(originalText3))
	//
	////4.4 使用rsa算法对数据进行签名
	//signText3, err := rsa.RSASign(pri, []byte(data3))
	//if err !=nil{
	//	fmt.Println("rsa算法签名失败",err.Error())
	//	return
	//}
	//
	////4.5 使用rsa公钥对签名进行验证
	//
	//verifyResult, err := rsa.RSAVerify(pri.PublicKey, []byte(data3), signText3)
	//if err != nil {
	//	fmt.Println("rsa签名验证失败：",err.Error())
	//	return
	//}
	//if verifyResult{
	//	fmt.Println("恭喜，rsa签名验证成功！")
	//}else {
	//	fmt.Println("抱歉，rsa签名验证失败！")
	//}


	//rsa.GeneratePriPem(pri)

	//5、ecc算法种的ecdsa数据签名算法
	priKey, err := ecc.GenerateKey()
	if err != nil {
		fmt.Println("ecdsa生成密钥错误：",err.Error())
		return
	}
	data5 := "我对鸡蛋过敏！"
	r, s, err := ecc.ECDSASign(priKey, []byte(data5))
	if err != nil {
		fmt.Println("签名错误：",err.Error())
		return
	}
	verifyResult := ecc.ECDSAVerify(priKey.PublicKey, r, s, []byte(data5))
	if verifyResult {
		fmt.Println("ecc签名验证成功")
	}else {
		fmt.Println("ecc签名验证失败")
	}
}
