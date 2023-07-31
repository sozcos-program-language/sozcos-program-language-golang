package main

import (
	"fmt"
)

/**
 * Go 语言变量的定义
 *
 * @author Elsen pooozg@gmail.com
 */
func main() {
	// 1.单行声明变量
	var v1 int
	s := "hello word"
	fmt.Println("声明变量: v1 =", v1, s)

	// 2.声明多个变量
	var (
		id   int8
		name string
		age  int
	)
	id = 1
	name = "小明"
	age = 10
	fmt.Println("id =", id, " name =", name, " age =", age)

	// 3.使用 = 声明, 编译器会根据 = 右边的类型进行推断, 但是这种方式仅限于函数内部使用
	name1 := "abc"
	fmt.Println(" name1 =", name1)
}
