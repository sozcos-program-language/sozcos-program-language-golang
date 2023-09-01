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
}
