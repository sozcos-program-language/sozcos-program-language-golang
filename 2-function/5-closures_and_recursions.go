/*
闭包与递归:

	闭包:函数引用外部变量, 闭包是一种函数概念, 使用场景较少, 个人认为可以理解成switch, 但是闭包是可以重复使用外部变量
*/
package main

import (
	"fmt"
)

// func main() {
// 	fmt.Printf(" %d\n", recursions(10))
//
// 	var i int
// 	for i = 0; i < 10; i++ {
// 		fmt.Printf("%d\n", aficionado(i))
// 	}
// }

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main1() {
	closuresCase1()
	closuresCase1()

	c := a()
	c()
	c()
	c()

	a() //不会输出i
}

// 闭包
func closuresCase1() int {
	num := 10

	closures := func() int {
		num++
		fmt.Println("num: >>>>>>>>>>>> num =", num)
		return num
	}()

	fmt.Println("closures: >>>>>>>>>>>> closures =", closures)
	return closures
}

// 递归
func recursions(i int) int {
	if i <= 1 {
		return 1
	}
	return i * recursions(i-1)
}

// 斐波那契数列
func aficionado(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return aficionado(i-1) + aficionado(i-2)
}
