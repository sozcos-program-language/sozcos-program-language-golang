/*
5-custom_error.go: 自定义异常
*/
package main

import (
	"fmt"
)

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
