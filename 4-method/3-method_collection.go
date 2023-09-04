/*
3-method_collection.go: 方法集——在go语言中，每个类型都有与之相关的方法，把这个类型的所有方法称为类型的方法集

	方法是为类型服务的， 而方法集则是该类型所有方法的集合。
*/
package main

import (
	"fmt"
)

type MethodCollection struct {
	int
}

func (mc MethodCollection) ToString() {
	fmt.Println("我是MethodCollection的ToString方法")
}

func (mc *MethodCollection) setName() {
	fmt.Println("我是MethodCollection的setName方法")
}
