/*
package main, 这是标记此项目是一个可执行的程序, 而不是库.
1.在使用了 package main 声明程序的时候, 程序必须有一个main函数入口,这个main函数在同一目录唯一.
2.每个 go 文件都可以使用 package main 来声明程序
3.同一个目录下的go文件中 package [包名] 必须一致
4.当前目录下所有文件使用了 package main 而没有 main 方法的时候, 编译会报错.

	执行: go build 命令返回以下错误:
	# go-practice/1-basic
	runtime.main_main·f: function main is undeclared in the main package

5.当同一目录内有多个main方法时, 编译会提示以下错误:

	.\2-package.go:7:6: main redeclared in this block
	.\1-package.go:23:6: other declaration of main

6.同理，当包的声明非 package main 或者 package init 的时候，其他的包声明都认为是库，库中的函数都是可被引用的
*/
package main

import (
	"fmt"
)

func main() {
	lens()
	lens2()
}

func lens() {
	fmt.Println("我是 lens 方法")
}
