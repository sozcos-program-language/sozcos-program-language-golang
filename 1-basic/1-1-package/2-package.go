package main

import (
	"fmt"

	myLib "go-practice/1-basic/1-2-lib"
)

func lens2() {
	fmt.Println("我是 lens2 方法")
	a := myLib.MyFunction{Name: "小明", Age: 10}
	fmt.Println(a)
}
