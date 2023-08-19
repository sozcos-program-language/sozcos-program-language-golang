/*
函数参数

	值传递:指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
	引用传递:是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

@author Elsen pooozg@gmail.com
*/
package main

import (
	"fmt"
)

// func main() {
// 	functionParamCase2("abc", 12, 10.0)
// 	s := []int{1, 2, 3}
// 	res := functionParamCase3("sum: %d", s...) // slice... 展开slice
// 	println(res)
// }

func functionParamCase1(x, y int, str string) int {
	fmt.Println(str)
	return x + y
}

// 使用 interface() 表示传递任意类型
func functionParamCase2(args ...interface{}) {
	for _, v := range args {
		fmt.Printf(" %v\n", v)
	}
}

func functionParamCase3(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}

	return fmt.Sprintf(s, x)
}
