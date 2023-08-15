/**
 * 1-for.go: for 初始化语句; 条件语句; 修饰语句 {}
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
)

func main() {
	// 1.基于计数器的迭代
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d\n ", i)
	}

	// Go 语言的平行赋值特性,可以使用多个计数器
	for i, j := 0, 12; i < j; i, j = i+1, j-1 {
		fmt.Printf("i = %d; j = %d\n", i, j)
	}

	// 2.基于条件判断的循环, 类似 Java while
	i := 5
	for i >= 0 {
		fmt.Println("我是for循环2: ", i)
		i -= 1
	}

	// 3.无限循环
	/*for {
		fmt.Println("我是无限循环")
		i += 1
	}*/

	// 4.k,v 循环
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	for k, v := range m {
		fmt.Println(k, v)
	}

	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("%c = %d\n", char, pos)
	}
}
