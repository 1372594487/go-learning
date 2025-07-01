/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-01 14:14:54
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-01 14:48:56
 * @FilePath: \go-learning\day1\func\func.go
 * @Description:函数多返回值写法
 *
 */
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// 多个返回值的函数,返回值类型用括号括起来
func swap(a, b string) (string, string) {
	return b, a
}

//匿名返回值 ,返回值类型可以省略

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 函数的参数类型可以省略,编译器会自动推断
func add2(a, b int) int {
	return a + b
}

// 函数的参数可以命名,返回值也可以命名
func split2(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 函数的参数可以命名,返回值也可以命名,返回值可以省略
func split3(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
	fmt.Println(add2(42, 13))
	fmt.Println(split2(17))
	fmt.Println(split3(17))
}
