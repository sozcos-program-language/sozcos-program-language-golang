/*
当包内的文件作为库的时候，可以让外部包引用
而包的引用可以使用别名，例如：

	import "abc/ddd/myLib" 改为 import myLib2 "abc/ddd/myLib"

在调用库的时候就可以使用别名作为前缀进行调用
*/

package mylib

func myListTest() {

}
