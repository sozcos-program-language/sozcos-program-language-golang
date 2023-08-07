package main

import (
	"fmt"
)

/**
 * 指针
 *
 * @author Elsen pooozg@gmail.com
 */
func main() {
	// go 语言中的指针是安全指针（不能进行偏移和运算），函数传参都是值拷贝，go的指针操作只有2中：&（取地址）和 *（根据地址取值）
	a := 1
	b := &a
	c := *b
	fmt.Println(" a =", a, ", b =", b, ", c =", c)

	// TODO new 以及 make：https://www.kancloud.cn/gofor/golang-learn/2571655
	// https://www.kancloud.cn/gofor/golang-learn/2120502
}
