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

	// 4.声明多个变量
	name, age = "小明", 31
	fmt.Println("age = ", age, " name =", name)
	// 变量转换
	var a int = 20
	var b int = 50
	b, a = a, b

	// 5.指针变量 & 普通变量：普通变量存储的是数据本身，指针变量存储的是地址
	// 普通变量
	var age1 int = 28
	var ptr = &age // &后面接变量名，表示取出该变量的内存地址
	fmt.Println("age1: ", age1)
	fmt.Println("ptr: ", ptr)
	// 指针变量——使用表达式 new(Type) 将创建一个Type类型的匿名变量，初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type。
	ptr2 := new(int)
	fmt.Println("ptr address: ", ptr2)
	fmt.Println("ptr value: ", *ptr2) // * 后面接指针变量，表示从内存地址中取出值
	// new函数类似是一种语法糖，而不是一个新的基础概念。不管哪种方法，变量/常量都只能声明一次，声明多次，编译就会报错
	/*
		// 使用 new
		func newInt() *int {
			return new(int)
		}

		// 使用传统的方式
		func newInt() *int {
			var dummy int
			return &dummy
		}
	*/
}
