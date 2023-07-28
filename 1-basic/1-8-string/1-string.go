package main

import (
	"fmt"
	"strings"
)

/**
 * Go 的字符串操作
 *
 * @author Elsen pooozg@gmail.com
 */
func main() {
	// 使用反引号 `(键盘1 旁边的) 定义换行字符串, 所有的转义字符均无效，文本将会原样输出
	str1 := `a
b
c`
	fmt.Println(str1)

	fmt.Println()
	str2 := `a`
	for i := range str2 {
		fmt.Println(i)
	}

	fmt.Println()
	str3 := "我是字符串\n我是字符串2"
	fmt.Println(str3)
	fmt.Println()

	/************ 字符串API **************/
	// len(str)	求长度
	s1 := "中/国/农/业/银/行"
	fmt.Println("len() == ", len(s1))
	// +或fmt.Sprintf	拼接字符串
	fmt.Println(s1 + "广东分行")
	// strings.Split	分割\
	fmt.Println(strings.Split(s1, "/"))
	for i := range strings.Split(s1, "/") {
		fmt.Println(i)
	}
	// strings.Contains	判断是否包含
	fmt.Printf("s1abc: ", s1)
	fmt.Println("abc s1", s1)
	fmt.Println(strings.Contains(s1, "/"))
	// strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	// strings.Index(),strings.LastIndex()	子串出现的位置
	// strings.Join(a[]string, sep string)	join操作
}
