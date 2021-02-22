package gamedb_open_api

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

func Get_Sha1(p string) string {

	data := []byte(p)
	txt := fmt.Sprintf("%x", sha1.Sum(data))
	//
	fmt.Printf("%x\n", sha1.Sum(data)) //597f6a540010f94c15d71806a99a2c8710e747bd
	//
	return txt
}

//参数转小写
func Get_Sha1_ParameterToLower(p string) string {

	var name = strings.ToLower(p)

	data := []byte(name)
	txt := fmt.Sprintf("%x", sha1.Sum(data))
	//
	fmt.Printf("%x\n", sha1.Sum(data)) //597f6a540010f94c15d71806a99a2c8710e747bd
	//
	return txt
}

//参数转小写
func Get_Md5_ParameterToLower(p string) string {

	var str = strings.ToLower(p)

	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

//参数转小写
func Get_Md5(p string) string {

	var str = strings.ToLower(p)

	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

//生成32位md5字串
func GetMd5String(s string) string {

	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetMD5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
