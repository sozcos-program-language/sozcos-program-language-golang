/**
 * 1-date_and_time.go: Go 的日期和时间 API
 *
 * @author Elsen pooozg@gmail.com
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前时间
	t := time.Now()
	fmt.Println("通用世界时间: >>>>>>>>>>>> t.UTC() =", t.UTC())
	fmt.Println("当前时间: >>>>>>>>>>>> time.Now() =", t)
	fmt.Println("当前日: >>>>>>>>>>>> time.Now().Day() =", t.Day())
	fmt.Println("当前分钟: >>>>>>>>>>>> time.Now().Minute() =", t.Minute())

	// 格式化时间, Go 的格式化时间只能使用 2006/01/02 15:04:05 来代替: yyyy/MM/dd HH:mm:ss
	// Printf 允许指定格式化字符串
	// %02d: 这是一个占位符，用于格式化整数值。%d 表示要插入一个整数值，而 02 表示要用零填充，确保输出至少两位数。如果数值不足两位，将在前面填充零。
	// %02d: 同样的占位符，用于格式化月份的整数值，确保输出至少两位数，并用零填充。
	// %4d: 这是一个占位符，用于格式化年份的整数值。%d 表示整数值，而 4 表示输出的最小宽度为 4 位，不足的地方将会用空格填充。
	fmt.Printf("日期格式化: >>>>>>>>>>>> %02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	fmt.Println("日期格式化: >>>>>>>>>>>> t.Format(\"2006 - 01 - 0215:04:05\") =", t.Format("2006-01-02 15:04:05"))

	// 时间操作API
	now := time.Now()
	// 创建指定时间
	specificTime := time.Date(2023, time.August, 11, 12, 30, 0, 0, time.UTC)
	fmt.Println("创建指定时间: >>>>>>>>>>>> time.Date(2023, time.August, 11, 12, 30, 0, 0, time.UTC) =", specificTime)

	// 获取时间的年、月、日、小时、分钟、秒、纳秒
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	nanosecond := now.Nanosecond()
	fmt.Printf("获取时间的年、月、日、小时、分钟、秒、纳秒 >>>>>>>>>>>> Year: %d, Month: %s, Day: %d, Hour: %d, Minute: %d, Second: %d, Nanosecond: %d\n",
		year, month, day, hour, minute, second, nanosecond)

	// 添加时间间隔
	oneHourLater := now.Add(time.Hour)
	fmt.Println("添加时间间隔 >>>>>>>>>>>> now.Add(time.Hour) = ", oneHourLater)

	// 计算两个时间的差距
	duration := oneHourLater.Sub(now)
	fmt.Println("计算两个时间的差距 >>>>>>>>>>>> ", duration)

	// 解析字符串为时间
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", "2023-08-11 10:30:00")
	fmt.Println("解析字符串为时间 >>>>>>>>>>>> ", parsedTime)

	// 判断一个时间是否在另一个时间之前、之后或相同
	fmt.Println("A 时间在 B 时间之前 >>>>>>>>>>>> specificTime.Before(now)", specificTime.Before(now))
	fmt.Println("A 时间在 B 时间之后 >>>>>>>>>>>> specificTime.After(now)", specificTime.After(now))
	fmt.Println("A 时间于 B 时间相同 >>>>>>>>>>>> specificTime.Equal(now)", specificTime.Equal(now))
}
