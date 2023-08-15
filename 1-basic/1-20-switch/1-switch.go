/**
 * 1-switch.go: switch 控制语句
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
)

func main() {
	switchStructure("1", 5)
	switchStructure2(1)
}

// switch 语句结构
func switchStructure(expr any, expr2 any) {
	// expr 可以是任何类型的参数, 比较值可以是同类型的任意值, 不局限于常量或整数
	switch expr {
	case "1":
		fmt.Println("case 1")
	case "2":
		fmt.Println("case 2")
	// 多个比较值
	case "3", "4":
		fmt.Println("case 3, 4")
	default:
		fmt.Println("我是默认值")
	}

	// 当一个函数有2个switch代码块并且都是用同一个变量进行比较的时候, 可以将他们当作是一个完整switch代码块, 但是 default 中的语句会被执行
	switch expr {
	case "5":
		fmt.Println("case 5")
	case "6":
		fmt.Println("case 6")
	default:
		fmt.Println("我是默认值2")
	}

	// 对于 expr, expr2, 可以使用有返回值的函数, case 除了采用常量值外, 也可以使用返回 bool 类型的表达式或者函数
	switch expr2 {
	case 5:
		fmt.Println("case 5")
	case 6:
		fmt.Println("case 6")
	default:
		fmt.Println("我是默认值3")
	}

	fmt.Println("expr =", expr, "expr2 =", expr2)
}

/*
每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。（ Go 语言使用快速的查找算法来测试 switch 条件与 case 分支的匹配情况，
直到算法匹配到某个 case 或者进入 default 条件为止。）
一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 break 语句来表示结束。
因此，程序也不会自动地去执行下一个分支的代码。如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的。
*/
func switchStructure2(expr any) {
	switch expr {
	case 1:
		fmt.Println("switch 2.1")
		fallthrough
	case 2:
		// 当 expr == 1 时也会执行
		fmt.Println("switch 2.2")
	}
}
