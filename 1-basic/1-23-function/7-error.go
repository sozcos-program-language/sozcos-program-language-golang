/*
7-error.go: 异常处理
panic: 抛出异常
recover: 捕获异常

使用场景: 抛出一个panic异常, 然后再defer中通过recover捕获这个异常
如何区别使用 panic 和 error 两种方式?
惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error。

注意:

	1.recover 捕获异常必须配合 defer 使用, 并且要在 panic 抛出异常之前定义好,利用recover处理panic指令，defer 必须放在 panic 之前定义，
	  另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
	2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
	3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。
*/
package main

import (
	"errors"
	"fmt"
)

/*
程序执行流程:
1.defer是在最后一个return的时候才执行, 对于没有return的函数,
*/
func errorCase1() {

	defer func() {
		fmt.Println("进入defer函数")
		if err := recover(); err != nil {
			fmt.Println("进入到defer函数")
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()

	// 当以下代码备注是, recover 函数则不会执行
	panic("errorCase1 error!")
}

/*
	按照 defer 先进后出的原色, 下面这段代码, 即使在 panic 抛出异常之后, defer 延迟函数还是会执行, 首先会执行 第二段的 defer, defer 抛出异

常, 最后再执行第一段 defer, 这里捕获异常只能捕获最近一次的异常.因此输出的是 defer panic
*/
func errorCase2() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("errorCase2 panic")
}

/*
捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递
*/
func errorCase3() {
	e1 := recover() // 无效
	fmt.Println(": >>>>>>>>>>>> e1 =", e1)

	defer func() {
		fmt.Println("errorCase3 defer")
		if e2 := recover(); e2 != nil {
			fmt.Println(": >>>>>>>>>>>> e2 =", e2) // 有效
		}
	}()

	defer except() // 有效

	panic("errorCase3 error")
}

func except() {
	fmt.Println(": >>>>>>>>>>>> except =", except)
	fmt.Println(recover()) // 捕获到异常之后不会再抛出异常, 其他 recover 无法再捕获
}

// MyError 标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型。
var MyError = errors.New("自定义异常")

func errorCase4() {
	defer func() {
		fmt.Println("捕获异常:")
		fmt.Println(recover())
	}()

	switch z, _, err := div(10, 1); err {
	case nil:
		println(z)
	case MyError:
		panic(err)
	}
}

func div(x, y int) (int, string, error) {
	if y == 0 {
		return 0, "", MyError
	}
	return x / y, "", nil
}

// Try 实现类似 try catch
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func errorCase5() {
	Try(func() {
		panic("errorCase5 panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}
