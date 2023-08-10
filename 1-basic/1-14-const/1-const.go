// 常量
//
// @author Elsen pooozg@gmail.com
package main

import (
	"fmt"
)

// 常量的定义
const name = "abc"

// 常量的值在编译期间就回确定,
const (
	A = 5
	B = "const value"
)

func main() {
	fmt.Println(" name =", name, " A =", A, " B =", B)
}
