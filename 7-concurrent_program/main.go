/*
main.go:
TODO: https://www.topgoer.com/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/runtime%E5%8C%85.html
其他板块未学习，
*/
package main

/*
程序启动会创建一个 goroutine，当main函数返回的时候， main函数中启动的goroutine会一同结束，
*/
func main() {
	// runGoroutine()
	runMultiple()
}
