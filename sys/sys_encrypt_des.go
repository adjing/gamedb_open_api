package sys

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

//des的CBC加密
//编写填充函数,如果最后一个分组字节不够,填充
//...字节数刚好合适,添加一个新分组
//填充个的字节的值==缺少的字节数
func paddingLastGroup(plainText []byte, bloclsize int) []byte {
	//1. 求出最后一个组中剩余的字节数
	padNum := bloclsize - (len(plainText) % bloclsize)
	//2. 创建一个新的切片, 长度==padNum, 每个字节值byte(padNum)
	char := []byte{byte(padNum)} //切片长度是1
	//切片创建并且重复多少次
	newPlain := bytes.Repeat(char, padNum)
	//3. newPlain数组追加到原始明文的后面
	newText := append(plainText, newPlain...)
	return newText
}

//去掉填充的数据
func unPaddingLastGrooup(plainText []byte) []byte {
	//1. 拿去切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	number := int(lastChar) //尾部填充的字节数
	return plainText[:length-number]
}

//des加密
func desEncrypt(plainText, key []byte) []byte {
	//1. 建一个底层使用des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2. 明文填充
	newText := paddingLastGroup(plainText, block.BlockSize())
	//3. 创建一个使用cbc分组接口
	iv := []byte("12345678") //8字节
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//4. 加密
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)
	//也可以这样,他加密过会把newText的值覆盖过去,然后返回newText就可以
	//blockMode.CryptBlocks(newText, newText)
	return cipherText
}

// des解密
func DesDecrypt(cipherText, key []byte) []byte {
	// 1. 建一个底层使用des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用cbc模式解密的接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3. 解密
	blockMode.CryptBlocks(cipherText, cipherText)
	// 4. cipherText现在存储的是明文, 需要删除加密时候填充的尾部数据
	plainText := unPaddingLastGrooup(cipherText)

	return plainText
}

// aes加密, 分组模式ctr
func aesEncrypt(plainText, key []byte) []byte {
	// 1. 建一个底层使用aes的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用ctr分组接口
	iv := []byte("12345678WHZdefgh")
	stream := cipher.NewCTR(block, iv)

	// 4. 加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

// des解密
func aesDecrypt(cipherText, key []byte) []byte {
	// 1. 建一个底层使用des的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用ctr模式解密的接口
	iv := []byte("12345678WHZdefgh")
	stream := cipher.NewCTR(block, iv)
	// 3. 解密
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText
}

//测试文件
func TestDES() {
	fmt.Println("des 加解密")
	key := []byte("1234abEd")
	src := []byte("特点: 密文没有规律,  明文分组是和一个数据流进行的按位异或操作, 最终生成了密文")
	cipherText := desEncrypt(src, key)
	plainText := DesDecrypt(cipherText, key)
	fmt.Printf("解密之后的数据: %s\n", string(plainText))

	fmt.Println("aes 加解密 ctr模式 ... ")
	key1 := []byte("1234abdd12345678")
	cipherText = aesEncrypt(src, key1)
	plainText = aesDecrypt(cipherText, key1)
	fmt.Printf("解密之后的数据: %s\n", string(plainText))
}

/*
func MyDesEncrypt(orig, key string) string{

    // 将加密内容和秘钥转成字节数组

    origData := []byte(orig)

    k := []byte(key)

    // 秘钥分组

    block, _ := des.NewCipher(k)

    //将明文按秘钥的长度做补全操作

    origData = PKCS5Padding(origData, block.BlockSize())

    //设置加密方式－CBC

    blockMode := cipher.NewCBCDecrypter(block, k)

    //创建明文长度的字节数组

    crypted := make([]byte, len(origData))

    //加密明文

    blockMode.CryptBlocks(crypted, origData)

    //将字节数组转换成字符串，base64编码

    return base64.StdEncoding.EncodeToString(crypted)

}


//DES解密方法
 func MyDESDecrypt(data string, key string) string {

    k := []byte(key)

    //将加密字符串用base64转换成字节数组

    crypted, _ := base64.StdEncoding.DecodeString(data)

    //将字节秘钥转换成block快

    block, _ := des.NewCipher(k)

    //设置解密方式－CBC

    blockMode := cipher.NewCBCEncrypter(block, k)

    //创建密文大小的数组变量
    origData := make([]byte, len(crypted))

    //解密密文到数组origData中

    blockMode.CryptBlocks(origData, crypted)

    //去掉加密时补全的部分

    origData = PKCS5UnPadding(origData)

    return string(origData)

}
holdtom
链接：https://www.imooc.com/article/271272/


*/
