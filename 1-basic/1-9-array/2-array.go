package main

import (
	"fmt"
)

/**
 * 切片(动态数组)
 *
 * @author Elsen pooozg@gmail.com
 */

func MySlices() {
	/*
		切片, 既动态数组, 可以通过长度和容量来追加元素
		切片中有两个概念：
			len长度，cap容量
			- 长度是指已经被赋过值的最大下标+1，可通过内置函数len()获得。
			- 容量是指切片目前可容纳的最多元素个数，可通过内置函数cap()获得
		切片是引用类型，因此在当传递切片时将引用同一指针，修改值将会影响其他的对象
	*/
	objs := []int{1, 2, 3}
	fmt.Println("切片数组: objs = ", objs)

	// 添加元素
	fmt.Println("添加元素 append(objs, 4) = ", append(objs, 4))
}
