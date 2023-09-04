/*
main.go:
*/
package main

import (
	"fmt"
)

func main() {
	// 方法
	user := Student{"Elsen", "高三1班"}
	user.Notify()

	// 匿名字段
	obj3 := Obj3{Obj1{"obj1_name"}, Obj2{"obj2_email"}}
	fmt.Println(obj3)
	fmt.Println("匿名字段: >>>>>>>>>>>> obj3.name =", obj3.name)
	fmt.Println("匿名字段: >>>>>>>>>>>> obj3.email =", obj3.email)
	fmt.Println(obj3.ToString())
	fmt.Println(obj3.Obj1.ToString())

	// 方法集
	m := MethodCollection{1}
	m.ToString()
	m.setName()

	// 表达式
	u := User1{1, "Tom"}
	u.Test()
	b := u.Test
	b() // 隐式传递 receiver
	a := (*User1).Test
	a(&u) // 显式传递 receiver

	// 自定义异常
	err := UpdateDatabase(0)
	switch v := err.(type) {
	case *CustomError:
		fmt.Println("get path error,", v)
	default:

	}
}
