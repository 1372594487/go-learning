/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-01 08:23:43
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-01 13:37:26
 * @FilePath: \go-learning\day1\var.go
 * @Description:变量声明
 *
 * Copyright (c) 2025 by ${git_name_email}, All Rights Reserved.
 */
package main

import "fmt"

// import (
// 	"fmt"
// 	"time"
// )

/* 变量声明 */
func main() {

	// 定义一个整型变量a，并赋值为10
	var a int = 10
	// 定义一个浮点型变量b，并赋值为20.5 短变量声明操作符:= 不支持全局变量声明
	b := 20.5
	// 定义一个字符串型变量c，并赋值为"start"
	var c = "start"

	//多变量声明
	// 定义两个整型变量d和e，并分别赋值为5和6
	var d, e int = 5, 6
	// 定义三个变量f、g和h，并分别赋值为1、2和false
	var f, g, h = 1, 2, false
	// 定义两个变量i和j，并分别赋值为"hello"和true
	var (
		i string = "hello"
		j        = true
	)

	// 将a转换为浮点型，并加上b，得到sum
	sum := float64(a) + b
	// 打印a、b、c、d、e、f、g、h、i、j和sum的值
	fmt.Println(a, b, c, d, e, f, g, h, i, j, sum)

}
