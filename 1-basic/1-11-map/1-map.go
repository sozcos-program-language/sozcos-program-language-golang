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
	mapOperator(scores)
}

/*
 *	Map 的遍历
 */
func traversalMap(m map[string]int) {
	fmt.Println()
	fmt.Println("***************** map 的遍历 *****************")
	// 1.带value
	for k, v := range m {
		fmt.Println("带value: ", " k =", k, "v =", v)
	}

	// 2.不带value
	for k := range m {
		fmt.Println("不带value: k =", k)
	}

	// 3.只需读取value，使用占位符, %d: 输出十进制的整数
	for _, score := range m {
		fmt.Printf("value: %d\n", score)
	}
}

/*
 * Map 的相关操作
 */
func mapOperator(m map[string]int) {
	fmt.Println()
	fmt.Println("***************** map 的相关操作 *****************")
	// 1.添加元素
	m["key2"] = 13

	for k, v := range m {
		fmt.Println(" k =", k, " v =", v)
	}

	// 2.获取指定key的值，当key不存在，会返回value类型的零值
	fmt.Println(" m[key1] =", m["key1"])
	fmt.Println(" m[key2] =", m["key2"])

	// 3.删除元素，如果key不存在，会静默处理
	delete(m, "key1")
	delete(m, "key2")
	fmt.Println(" m =", m)

	// 4.判断key是否存在：字典的下标读取可以返回两个值，使用第二个返回值都表示对应的 key 是否存在，若存在ok为true，若不存在，则ok为false
	scores := map[string]int{"english": 80, "chinese": 85}
	math, ok := scores["math"]
	if ok {
		fmt.Printf("math 的值是: %d", math)
	} else {
		fmt.Println("math 不存在")
	}
}
