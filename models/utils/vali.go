package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// TODO: 验证码  <28-12-21, Redari-Es> //
/*
func Vali() {
	for i := 0; i < 6; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(128))
		fmt.Println(result)
	}
}

*/
func ValiCode(n int) interface{} {
	//str := strings.ToLower(randSeq(n))
	str := randSeq(n)
	fmt.Println(str)
	return str
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}
	return string(b)
}

//传入的数据不一样，那么MD5后的32位长度的数据也不同
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func Trim(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	return s
}
func Len(str string) bool {
	b := false
	if len(str) >= 8 {
		fmt.Println("符合字符长度")
		b = true
	} else {
		fmt.Println("不符合字符长度：", len(str), "位")
	}
	return b
}
func Uid(id int) string {
	uid := "100001" + strconv.Itoa(id)
	return uid
}
