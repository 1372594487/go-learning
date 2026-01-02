package main

import (
	"fmt"
	"time"
)

const (
	//定义时间格式
	timeFormat = "2006-01-02 15:04"
)

func main() {
	nowTime := time.Now()
	fmt.Println(nowTime)
	//2023-03-28 10:53:39.843925 +0800 CST m=+0.000060874

	//加几分钟
	addMinutes := nowTime.Add(time.Minute * 10)
	fmt.Println(addMinutes)

	//减5天
	subtractDays := nowTime.AddDate(0, 0, -5)
	fmt.Println(subtractDays)

	//获取两个日期的天数(时间差)
	days := subtractDays.Sub(addMinutes).Hours() / 24
	fmt.Println(days)

	//哪个时间更大
	fmt.Println(nowTime.After(addMinutes))
	fmt.Println(nowTime.Before(addMinutes))

	fmt.Println(nowTime.Format(time.DateTime)) //YYYY-MM-DD HH:MM:SS

	loc, _ := time.LoadLocation("America/New_York")
	nyTime := time.Now().In(loc)
	fmt.Println(nyTime.Format(time.DateTime))
	fmt.Println(nyTime.Format(timeFormat))

	dtime := "2025-10-18 14:23:12"
	ddtime, _ := time.Parse(time.DateTime, dtime)
	fmt.Println(ddtime.Format(time.DateTime))

	//时间周期
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2025, 12, 31, 0, 0, 0, 0, time.Local)
	days = end.Sub(start).Hours() / 24
	fmt.Println(days)
}
