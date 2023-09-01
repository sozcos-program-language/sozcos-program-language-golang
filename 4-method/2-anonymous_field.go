/*
2-anonymous_field.go:	匿名字段,类似于Java中对象包含对象的用法, 也类似于继承
*/
package main

import (
	"fmt"
)

type Obj1 struct {
	name string
}

type Obj2 struct {
	email string
}

type Obj3 struct {
	Obj1
	Obj2
}

// ToString 不同对象采用相同的方法名, 就可以实现重载
func (self *Obj1) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Obj3) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}
