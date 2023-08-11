/**
 * 1-strings.go Go 语言字符串操作函数
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
	"strings"
)

var str = "Maybe let something else replace me in this world."

func main() {
	fmt.Println("判断字符串是否以 xx 开头: strings.HasPrefix(str, \"m\") =", strings.HasPrefix(str, "M"))
	fmt.Println("判断字符串是否以 xx 结尾: strings.HasSuffix(str, \"d\") =", strings.HasSuffix(str, "d"))
	fmt.Println("判断字符串是否包含 xx: strings.Contains(str, z) =", strings.Contains(str, "z"))
}
