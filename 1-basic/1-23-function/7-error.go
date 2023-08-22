/*
7-error.go: 异常处理
panic: 抛出异常
recover: 捕获异常
使用场景:

	抛出一个panic异常, 然后再defer中通过recover捕获这个异常

https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%BC%82%E5%B8%B8%E5%A4%84%E7%90%86.html
*/
package main

var _sbc string

func getSbc() string {
	return _sbc
}
