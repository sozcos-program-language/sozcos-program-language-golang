package main

import (
	"fmt"
)

func main() {
	fmt.Println("iota 的使用方式1:")
	fmt.Println("ConstValue1 = ", ConstValue1)
	fmt.Println("ConstValue2 = ", ConstValue2)
	fmt.Println("ConstValue3 = ", ConstValue3)
	fmt.Println("-----------------------------")
	fmt.Println("iota 的使用方式2:")
	fmt.Println("C21 = ", C21)
	fmt.Println("C22 = ", C22)
	fmt.Println("C23 = ", C23)
	fmt.Println("-----------------------------")
	fmt.Println("iota 的使用方式3:")
	fmt.Println("C31 = ", C31)
	fmt.Println("C32 = ", C32)
	fmt.Println("C33 = ", C33)
	fmt.Println("-----------------------------")
	fmt.Println("iota 的使用方式4:")
	fmt.Println("C41 = ", C41)
	fmt.Println("C42 = ", C42)
	fmt.Println("C43 = ", C43)
	fmt.Println("-----------------------------")
	fmt.Println("iota 的使用方式5:")
	fmt.Println("C51 = ", C51, "C54 = ", C54)
	fmt.Println("C52 = ", C52, "C55 = ", C55)
	fmt.Println("C53 = ", C53, "C56 = ", C56)
	fmt.Println("-----------------------------")
	fmt.Println("iota 的使用方式6:")
	fmt.Println("C61 = ", C61)
	fmt.Println("C62 = ", C62)
	fmt.Println("C63 = ", C63)
	fmt.Println("C64 = ", C64)
	fmt.Println("C65 = ", C65)
	fmt.Println("C66 = ", C66)
	fmt.Println("C67 = ", C67)
	fmt.Println("-----------------------------")
	fmt.Println("iota 的使用方式7:")
	fmt.Println("C71 = ", C71)
	fmt.Println("C72 = ", C72)
	fmt.Println("C73 = ", C73)
	fmt.Println("C74 = ", C74)
}

/*
1.iota: 常量计数器, 只能在常量的表达式中使用, 在const中出现的时候被重置为0, 可以理解成类似于 java 的 i++
2.iota的作用范围仅限于自已所在的const声明，多个const都使用iota会不会互相影响
3.如果常量没有表达式，会继承它上面常量的表达式。第一个常量必须有表达式，否则会编译失败 --> 示例6

iota 是如何进行计算?

	常量中有一个概念叫常量编号, 这个编号并不是常量的顺序号，而是const内部的行号，行号从0开始:
		const (
			A = iota // 0
			_ // 1
			B // 2
			C // 3
		)
	iota 的计算就是通过编号进行计算:
		const (
			A = iota
			B = 1 << iota
			C = -(iota)
			D = iota*10
		)

		常量A的表达式为：iota，编号为0，将编号替换表达式的iota，表达式为0，最后A的值为0。
		常量B的表达式为：1 << iota，编号为1，将编号替换表达式的iota，表达式为1<<1，B的值为2。
		常量C的表达式为：-(iota)，编号为2，将编号替换表达式的iota，表达式为-(2)，B的值为-2。
		常量D的表达式为：iota*10，编号为3，将编号替换表达式的iota，表达式为```3*10``，B的值为30。
*/
const (
	A = iota
	B
)

/* iota 的使用方式 1 : 给常量声明一个类型 */
const (
	ConstValue1 ParamValue = iota
	ConstValue2
	ConstValue3
)

type ParamValue int

/* iota 的使用方式 2 : 使用 "_" 跳过常量的 iota 赋值 */
const (
	C21 = iota
	_
	C22
	C23
)

/* iota 的使用方式 3 */
const (
	C31 = 1
	C32 = iota
	C33
)

/* iota 的使用方式 4 */
const (
	C41 = 1 << iota
	C42
	C43
)

/* iota 的使用方式 5 */
const (
	C51, C54 = iota + 1, iota + 2
	C52, C55
	C53, C56
)

/*
iota 的使用方式 6:

	常量B会继承常量A的表达式；常量D会继承常量C的表达式；常量F会继承常量E的表达式；
	常量G会继承常量F的表达式，由于F继承了E的，那相当于G也继承了E的表达式。
*/
const (
	C61 = iota
	C62
	_
	C63 = 3
	_
	C64
	C65 = iota << 1
	C66
	C67
)

/* iota 的使用方式 7 */
const (
	C71 = iota
	C72 = 100
	C73 = iota
	C74
)
