package main

import (
	"fmt"
)

/**
 * 数组
 *
 * @author Elsen pooozg@gmail.com
 */
func main() {
	// 1.声明一个由10个字符串组成的一维字符串数组, Go 语言中，新建的数组声明了大小后续便无法修改, 元素数量必须符合声明的数组长度
	var arr1 [10]string
	fmt.Println(" arr1 = ", arr1)

	// 2.数组的初始化
	var arr2 = [2]string{"a", "a"}
	// var arr2 = [2]string{"a", "a", "c"} // 超出数组长度提示: index 2 is out of bounds
	fmt.Println(" arr2 = ", arr2)

	// 3.可变数组
	arr01 := [...]int{1, 2, 3}
	arr02 := [...]int{1, 2, 3, 4}
	fmt.Printf("%d 的类型是: %T\n", arr01, arr01)
	fmt.Printf("%d 的类型是: %T", arr02, arr02)
	fmt.Println()
	// 二维数组, 二维数组的第二个数组长度不可忽略
	var twoArrs [3][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
	fmt.Println("二维数组的定义 twoArrs = ", twoArrs)

	// 4.定义数组别名类型: 每次定义数组都要书写: [3]int 会很麻烦, 因此可以将 [3]int 赋值给一个类型别名, 下次再次创建同类型长度的数组的时候可使用
	type intArrType [3]int
	ia1 := intArrType{1, 2, 3}
	fmt.Println()
	fmt.Printf("%d 的类型是: %T", ia1, ia1)

	// 5.数组参数, go 语言中的数组传递是通过值传递, 而不是地址传递
	// 声明并初始化一个整型数组
	intArr := [10]int{5: 10}
	// 打印这个整型数组的地址
	fmt.Println()
	fmt.Printf("In main function the address is : %p\n", &intArr)
	// 调用test函数f
	test(intArr)

	// 6.动态数组
	fmt.Println("\n****************** 动态数组 ******************")
	MySlices()
}

func test(param [10]int) {
	// 打印入参的整型数组的地址
	fmt.Printf("In test function the address is : %p\n", &param)
}
