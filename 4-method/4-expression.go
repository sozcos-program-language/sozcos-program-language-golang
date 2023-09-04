/*
4-expression.go: 表达式--TODO??
https://www.topgoer.com/%E6%96%B9%E6%B3%95/%E8%A1%A8%E8%BE%BE%E5%BC%8F.html
*/
package main

import (
	"fmt"
)

type User1 struct {
	id   int
	name string
}

func (self *User1) Test() {
	fmt.Printf("%p, %v\n", self, self)
}
