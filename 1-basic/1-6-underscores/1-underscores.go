/*
下划线：
go 语言中的下划线表示忽略
1.有些场景，我们只需要是引入的库去执行初始化函数，而不需要使用到库中的函数：

	import _ "go-practice/1-basic/1-2-lib"

2.代码中的下划线，表示忽略这个变量
*/
package main

import (
	"fmt"
	"os"

	_ "go-practice/1-basic/1-4-init"
)

func main() {
	buf := make([]byte, 1024)
	/*
		因为go语言中的机制——当变量没有被使用的时候会编译报错，因此，有些方法返回值是多个的时候，而我们也不需要使用其中返回的变量，我们就可以使
		用“_”来忽略，从而避免编译报错
		f, err := os.Open("filePath") --> f, _ := os.Open("filePath")
	*/
	f, _ := os.Open("go.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		fmt.Println(os.Stdout.Write(buf[:n]))
	}
}
