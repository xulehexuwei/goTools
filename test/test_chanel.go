package main

import "fmt"

func main() {
	//f := "f"

	fmt.Println("嗨客网(www.haicoder.net)")
	//使用字符串切片的形式，截取字符串
	s := "嗨客网"
	str := []rune(s)
	fmt.Println(len(str))
	str1 := str[0:1]
	str2 := str[:3]
	str3 := str[1:]
	fmt.Println("str1 =", string(str1), "str2 =", str2, "str3 =", str3)
}
