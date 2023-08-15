/**
 * 1-control_structures.go:
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ifStructures()
	switchStructures()
	forStructures()
}

// if 控制语句
func ifStructures() {
	// Go语言中没有三元运算符, 因此需要使用 if else的形式类实现,这体现了Go语言只用一种方法解决问题的设计
	/*
		if randNum := rand.Int(); randNum > 1 {
		等价于
		randNum := rand.Int();
		if randNum > 1 {
	*/
	if randNum := rand.Int(); randNum > 3801492532031179764 {
		fmt.Println(": >>>>>>>>>>>> a", randNum)
	} else {
		fmt.Println(": >>>>>>>>>>>> b", randNum)
	}
}

// switch 控制语句
func switchStructures() {
	randNum := 12
	str := ""
	switch {
	case randNum <= 4:
		str = "A"
	case randNum >= 5:
		str = "B"
	}
	fmt.Println(": >>>>>>>>>>>> str =", str)

	// switch后面可以带上一个初始化语句
	switch i := "y"; i {
	// 多个case值使用逗号分隔
	case "y", "Y":
		fmt.Println("yes")
		// fallthrough会跳过接下来的case表达式直接执行下一个case语句
		fallthrough
	case "n", "N":
		fmt.Println("no")
	}
}

// for 循环
func forStructures() {
	arr1 := [...]int{1, 2, 3, 4, 5}
	for par := range arr1 {
		fmt.Println(": >>>>>>>>>>>> par =", par)
	}

	// 声明一个新数组
	var m = make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2

	for k, v := range m {
		fmt.Println(k, v)
	}
}
