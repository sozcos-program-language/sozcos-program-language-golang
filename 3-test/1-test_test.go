/*
Golang 单元测试：

	1.使用 go test 命令执行单元测试
	2.单元测试文件必须以 _test.go 结尾, go test 命令会遍历所有的 *_test.go 文件中符合以下命名规则的函数，然后生成一个临时的main包用于调用相应
	  的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件.
	3._test.go 结尾的文件又有三种类型的函数:
	  a.单元测试函数: 函数名前缀为Test,测试程序的一些逻辑行为是否正确
	  b.基准测试函数: 函数名前缀为Benchmark,测试函数的性能
	  c.示例函数: 函数名前缀为Example,为文档提供示例文档

测试函数的基本格式:

	func TestName(t *testing.T){
		// ...
	}
*/
package main

import (
	"reflect"
	"strings"
	"testing"
)

/*
1、文件名必须以xx_test.go命名
2、方法必须是Test[^a-z]开头
3、方法参数必须 t *testing.T
4、使用go test执行单元测试
*/
func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
