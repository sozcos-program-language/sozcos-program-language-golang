/**
 * 1-strings.go Go 语言字符串操作函数
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var str = "|Maybe let something else replace me in this world.|"

func main() {
	fmt.Println("判断字符串是否以 xx 开头: >>>>>>>>>>>> strings.HasPrefix(str, m) =", strings.HasPrefix(str, "M"))
	fmt.Println("判断字符串是否以 xx 结尾: >>>>>>>>>>>> strings.HasSuffix(str, d) =", strings.HasSuffix(str, "d"))
	fmt.Println("判断字符串是否包含 xx 字符: >>>>>>>>>>>> strings.Contains(str, z) =", strings.Contains(str, "z"))
	fmt.Println("字符所在字符串位置, -1 表示不包含字符: >>>>>>>>>>>> strings.Index(str, in) =", strings.Index(str, "in"))
	fmt.Println("字符所在字符串最后一次出现的位置, -1 表示不包含字符: >>>>>>>>>>>> strings.LastIndex(str, d) =", strings.LastIndex(str, "d"))
	fmt.Println("字符串替换: -1 或超出原有字符数量则为全部替换: >>>>>>>>>>>> strings.Replace(str, Maybe, wait, 1) =", strings.Replace(str, "a", "啊", 3))
	fmt.Println("统计字符出现次数: >>>>>>>>>>>> strings.Count(str, e) =", strings.Count(str, "e"))
	fmt.Println("复制字符串: >>>>>>>>>>>> strings.Repeat(hello!, 3) =", strings.Repeat("hello!", 3))
	fmt.Println("转小写: >>>>>>>>>>>> strings.ToLower(str) =", strings.ToLower(str))
	fmt.Println("转大写: >>>>>>>>>>>> strings.ToUpper(str) =", strings.ToUpper(str))

	fmt.Println("修剪字符串前后空格: >>>>>>>>>>>> strings.TrimSpace(str) =", strings.TrimSpace(str))
	fmt.Println("修建字符串前后指定字符: >>>>>>>>>>>> strings.Trim(str, |) =", strings.Trim(str, "|"))
	fmt.Println("修建字符串开头字符: >>>>>>>>>>>> strings.TrimLeft(str, |) =", strings.TrimLeft(str, "|"))
	fmt.Println("修建字符串结尾字符: >>>>>>>>>>>> strings.TrimRight(str, |) =", strings.TrimRight(str, "|"))

	fmt.Println("根据空格分割字符串, 并返回 slice: >>>>>>>>>>>> strings.Fields(str) =", strings.Fields(str))
	fmt.Println("根据指定字符分割字符串: >>>>>>>>>>>> strings.Split(str, e) =", strings.Split(str, "e"))

	// 类型转换, 其他类型转为 string 总是成功的, 但是 string 转为其它类型并不保证成功
	num := 12
	fmt.Println("数字转字符串: >>>>>>>>>>>> strconv.Itoa(num) =", strconv.Itoa(num))
	//  将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。
	num2 := 18.891
	fmt.Println("浮点数转字符串: >>>>>>>>>>>> strconv.FormatFloat(num2, 'f', 5, 64) =", strconv.FormatFloat(num2, 'f', 5, 64))

	num3 := "32a"
	// 字符串转数字, 因为字符串转数字不一定成功,因此返回多结果,result是结果, err是判断是否成功
	result, err := strconv.Atoi(num3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("字符串转数字: >>>>>>>>>>>> result =", result)
}
