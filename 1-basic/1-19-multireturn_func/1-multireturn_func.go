/**
 * 1-multireturn_func.go: 多返回值函数
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// 根据函数返回值判断是否有异常
	result, err := strconv.Atoi("12")

	// 判断程序是否有err的方式1
	if err != nil {
		// 打印异常
		fmt.Println(": >>>>>>>>>>>> err =", err)
		// 终止程序运行
		os.Exit(1)
		return
	} else {
		fmt.Println("字符串转数字: >>>>>>>>>>>> result =", result)
	}

	// 判断程序是否有err的方式2
	if r, e1 := strconv.Atoi("14"); e1 == nil {
		fmt.Println("判断程序是否有err的方式2: >>>>>>>>>>>> r =", r)
	}

	// 判断程序是否有err的方式3
	if v, ok := mySqrt(-12); ok {
		fmt.Println("判断程序是否有err的方式3: >>>>>>>>>>>> v =", v)
	}

	// 使用下划线忽略异常
	atoi, _ := strconv.Atoi("14a2")
	fmt.Println("使用下划线忽略异常: >>>>>>>>>>>> atoi =", atoi)

}

// 定义一个多返回值函数, 根据执行结果进行判断
func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		// error case
		return
	}
	return math.Sqrt(f), true
}
