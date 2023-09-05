/*
1-interface.go: 接口——接口（interface）定义了一个对象的行为规范;

		GO 中，接口是一种类型, 一种抽象的类型，它是一组方法的集合（方法集？）
		接口定义格式：
			type 接口类型名 interface{
				方法名1( 参数列表1 ) 返回值列表1
				方法名2( 参数列表2 ) 返回值列表2
				…
			}
	    1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
	    2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
	    3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略

	Go 语言中, 一个类型可以实现多个接口,TODO
*/
package main

import (
	"fmt"
)

// Sayer 定义接口
type Sayer interface {
	Say()
	Eat()
}

// 声明结构体

type Cat struct {
}

type Dog struct {
}

// 实现接口

func (c Cat) Say() {
	fmt.Println("喵喵喵")
}

func (c Cat) Eat() {
	fmt.Println("喵喵喵Eat")
}

func (d Dog) Say() {
	fmt.Println("汪汪汪")
}

func (d Dog) Eat() {
	fmt.Println("汪汪汪Eat")
}

// 接口类型变量能够存储所有实现了该接口的实例
func runCase1() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := Cat{}  // 实例化一个cat
	b := Dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.Say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.Say()     // 汪汪汪
}
