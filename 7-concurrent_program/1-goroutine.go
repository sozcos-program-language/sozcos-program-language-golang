/*
1-goroutine.go:

	goroutine——程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行

goroutine类似线程，但是由Go调度和管理，go程序会智能地将 goroutine 中的任务合理的分配给每个
CPU, 因为Go在语言层面已经内置了调度和上下文切换的机制。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// 实现 goroutine 同步
var wg sync.WaitGroup

// 使用goroutine，只需要在调用函数的时候在前面加上 go 关键字，就可以为一个函数创建一个 goroutine
func case1() {
	fmt.Println("hello goroutine")
}

func runGoroutine() {
	go case1() // 启动一个线程执行该函数
	fmt.Println("main func")

	// 使用延迟函数，让goroutine创建并执行
	time.Sleep(time.Second)
}

// 运行多个goroutine
func runMultiple() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个 goroutine 就等级 +1
		go case2(i)
	}
	wg.Wait() // 等待所有登记的 goroutine 都结束
}

func case2(i int) {
	defer wg.Done() // goroutine 结束就 等级 -1
	fmt.Println("goroutine case2 >> ", i)
}
