package main

import (
	"fmt"
)

/**
 * Go 的 Map 类型
 *
 * @author Elsen pooozg@gmail.com
 */
func main() {
	// 1.声明 Map 方式1
	var map1 map[string]int
	// 只声明，不赋值的map，值为 nil，此时不能赋值，需要进行初始化 ： panic: assignment to entry in nil map
	if map1 == nil {
		map1 = make(map[string]int)
	}
	// 初始化后就可以进行赋值
	map1["key"] = 12
	fmt.Println(" map1 =", map1)

	// 2.声明 Map 方式2
	map2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(" map2 =", map2)

	// 3.声明 Map 方式3
	scores := make(map[string]int)
	scores["english"] = 80
	scores["chinese"] = 85

	traversalMap(scores)
}

/*
 *	Map 的遍历
 */
func traversalMap(m map[string]int) {
	// 1.带value
	for k, v := range m {
		fmt.Println("带value: ", " k =", k, "v =", v)
	}

	// 2.不带value
	for k := range m {
		fmt.Println("不带value: k =", k)
	}
}
