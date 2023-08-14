/**
 * multireturn_func.go:
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 根据函数返回值判断是否有异常
	result, err := strconv.Atoi("12a")

	if err != nil {
		// 打印异常
		fmt.Println(": >>>>>>>>>>>> err =", err)
		// 终止程序运行
		os.Exit(1)
		return
	}
	fmt.Println("字符串转数字: >>>>>>>>>>>> result =", result)
}
