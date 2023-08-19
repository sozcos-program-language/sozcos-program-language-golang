/*
延时调用 defer:
 1. 关键字 defer 用于注册延迟调用。
 2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
 3. 多个defer语句，按先进后出的方式执行。
 4. defer语句中的变量，在defer声明时就决定了。

用途:

	1.关闭文件句柄——即关闭程序的访问权限，以清除该程序的所有活动和资源。它可以防止程序运行中的错误，以及可能导致系统崩溃的错误操作。关闭句柄也可以保
			护用户的隐私，防止未经授权的程序访问用户数据。一个句柄就是你给一个文件，设备，套接字(socket)或管道的一个名字, 以便帮助你记住你
			正处理的名字, 并隐藏某些缓存等的复杂性, 文件句柄是一种用于跟踪和管理文件访问的抽象概念，使程序能够与操作系统交互地操作文件。
	2.锁资源释放
	3.数据库连接释放
*/
package main

import (
	"fmt"
)

// func main() {
//
// 	// defer 先进后出
// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
//
// 	// 中间的defer执行报错, 其他defer也会执行
// 	defer func() {
// 		func2(0)
// 	}()
//
// 	defer fmt.Println("3")
// 	defer fmt.Println("4")
//
// 	func1(10)
// }

func func2(i int) {
	// div0 异常未被捕获，逐步往外传递，最终终止进程。
	println(100 / i)
}

func func1(i int) int {
	fmt.Println("func1: >>>>>>>>>>>> i =", i)
	return i
}
