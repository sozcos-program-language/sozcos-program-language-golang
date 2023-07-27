package myInit

/*
init 函数是go自动调用, 当我们使用到init函数的时候需要注意它们的调用顺序
 1. 同个文件内, init 函数按照从上到下的顺序来调用
 2. 不同文件, 同一个包内的文件是按照文件名字符串比较"从小到大"顺序嗲用个文件中的init函数
 3. 对于不同包, 要分2中情况:
    3.1 文件不相互依赖, 按照main包中先import的后调用的顺序调用其保重的init
    3.2 如果存在依赖, 先调用最早被依赖的包中的init,
    也就是说,先被用到的文件的init函数会先调用
 4. 如果init函数使用了println()函数或者print(),执行顺序则对以上规则失效,因此正式环境中不能使用println()呵print()函数
*/
func init() {
	println("我是 init 方法1")
}

func init() {
	println("我是 init 方法2")
}
