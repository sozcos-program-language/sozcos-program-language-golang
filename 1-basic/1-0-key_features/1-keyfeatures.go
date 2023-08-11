// Go 程序的一般结构
//
// @author Elsen pooozg@gmail.com
package main

// 导包
import (
	"fmt"
)

// 常量
const c = "C"

// 全局变量
var v int = 5

// 类型定义
type T struct{}

// 初始化函数
func init() { // initialization of package
}

// main 函数
func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

// 自定义函数
func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	//...
}
