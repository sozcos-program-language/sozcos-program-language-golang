/*
1-constructor.go:	结构体
Go 语言中没有类的概念, 也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性
*/
package main

import (
	"fmt"
)

// MyInt 自定义类型: 通过 type 关键字创建一个类型, 也可以使用 struct, 通过Type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
type MyInt int

// NewInt 类型别名, 也就是相当于 var a = int, 一个别名
type NewInt = int

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("a ===> %T\n", a)
	fmt.Printf("b ===> %T\n", b)

	structCase1()
	structCase2()
	structCase3()
	structCase4()
}

// 声明一个结构体实例
func structCase1() {
	var member Member
	member.id = 1
	member.age = 18
	member.account = "user"
	member.name = "小明"
	fmt.Println(": >>>>>>>>>>>> member =", member)
}

// 匿名结构体
func structCase2() {
	var user struct {
		age  int
		name string
	}

	user.age = 1
	user.name = "匿名结构体"

	fmt.Println("匿名结构体: >>>>>>>>>>>> user =", user)
}

// Member 定义一个结构体, 类型名 包内唯一
type Member struct {
	id            int64
	age           int8
	account, name string
}

// 创建指针类型结构体
func structCase3() {
	// Go 支持通过地址访问结构听成员变量
	m := new(Member)
	fmt.Printf("%T\n", m)
	fmt.Println("创建指针类型结构体: >>>>>>>>>>>> m.age =", m.age)
}

// 结构体的初始化
func structCase4() {
	member := Member{
		1,
		12,
		"account",
		"name",
	}

	p := person{
		1,
		"name",
	}

	fmt.Println("结构体的初始化: >>>>>>>>>>>> Member =", member)
	fmt.Println("结构体的初始化: >>>>>>>>>>>> person =", p)
}

type person struct {
	id   int
	name string
}
