/**
 * 1-basic_types.go Go 的基本类型
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
)

// 布尔类型
var b bool

// 数字类型
var i1 int
var i2 int8
var i3 int16
var i4 int32
var i5 int64

// 无符号类型 >>> 大于等于0的整数
var ui1 uint
var ui2 uint8
var ui3 uint16
var ui4 uint32
var ui5 uint64

// 浮点型
// 精确到小数点后 7 位
var f1 float32

// 精确到小数点后 15 位, 尽可能地使用 float64，因为 math 包中所有有关数学运算的函数都会要求接收这个类型
var f2 float64

// 字符串, go 的编译器会自动追加分号,因此在拼接字符串的时候 + 必须在第一行
var s1 string

func main() {
	fmt.Printf(" b = %T\n", b)
	fmt.Println(" i1 =", i1, " i2 =", i2, " i3 =", i3, " i4 =", i4, " i5 =", i5)
	fmt.Println(" ui1 =", ui1, " ui2 =", ui2, " ui3 =", ui3, " ui4 =", ui4, " ui5 =", ui5)
	fmt.Println(" f1 =", f1, " f2 =", f2)

	s2 := s1 + "s" + "b"
	println(s2)
}
