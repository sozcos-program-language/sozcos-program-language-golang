/*
1-method_define.go: 方法定义

	方法与函数不同, 方法只能是为当前包内命名类型定义方法
	方法中的形参指向了可以使用该方法的类型
*/
package main

import (
	"fmt"
)

type User struct {
	name  string
	email string
}

type Student struct {
	name  string
	class string
}

func (s Student) Notify() {
	fmt.Println(s)
}
