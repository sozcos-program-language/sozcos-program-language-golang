/*
5-custom_error.go: 自定义异常
*/
package main

import (
	"errors"
	"fmt"
)

// 构建新异常
func getCircleArea(radius float32) (area float32, err error) {
	if radius < 0 {
		// 构建个异常对象
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14 * radius * radius
	return
}

// CustomError 自定义异常
type CustomError struct {
	table   string
	message string
}

func (p *CustomError) Error() string {
	return fmt.Sprintf("table_name=%s \nmessage=%s", p.table, p.message)
}

func UpdateDatabase(updateNum int) error {

	if updateNum <= 0 {
		return &CustomError{
			table:   "table",
			message: "自定义异常, 更新数据库数量为0",
		}
	}

	return nil
}
