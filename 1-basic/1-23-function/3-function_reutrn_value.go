/*
函数返回值
@author Elsen pooozg@gmail.com
*/
package main

// 函数的对各返回值可以直接作为其他函数的参数
func test1() (int, int) {
	return 1, 2
}

// 直接返回, 但只是应用在比较短的函数中
func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

func add(x, y int) int {
	return x + y
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}

	return x
}
