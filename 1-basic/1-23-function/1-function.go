/*
函数的定义:

@author Elsen pooozg@gmail.com
*/
package main

import (
	"fmt"
)

func main() {
	p1, p2 := functionCase1(1, 2, "as")
	fmt.Printf(" %d , %v", p1, p2)

	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println(s1, s2)
}

func functionCase1(x, y int, s string) (int, string) {
	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func test(fn func() int) int {
	return fn()
}

// FormatFunc 定义函数类型——函数可以声明为一个类型， 再使用的时候当作形参来传入
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
