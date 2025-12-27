/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-27 02:06:29
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-27 02:13:39
 * @Description: fmt
 */
package main

import "fmt"

func main() {
	name := "Alice"
	age := 25
	height := 1.65
	isStudent := true

	// %s - 字符串
	fmt.Printf("Name: %s\n", name)

	// %d - 十进制整数
	fmt.Printf("Age: %d\n", age)

	// %f - 浮点数
	fmt.Printf("Height: %.2f\n", height) // .2f 表示保留2位小数

	// %t - 布尔值
	fmt.Printf("Is student: %t\n", isStudent)

	// %v - 值的默认格式
	fmt.Printf("Default: %v\n", name)

	// %+v - 对于结构体，会显示字段名
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Bob", Age: 30}
	fmt.Printf("Struct with fields: %+v\n", p)

	// %#v - 值的Go语法表示
	fmt.Printf("Go syntax: %#v\n", p)

	// %T - 值的类型
	fmt.Printf("Type: %T\n", p)

	// %% - 字面上的百分号
	fmt.Printf("Progress: 80%%\n")

	// 宽度和精度
	s := "hello"
	fmt.Printf("|%s|\n", s)     // |hello|
	fmt.Printf("|%10s|\n", s)   // |     hello|  (右对齐，总宽度10)
	fmt.Printf("|%-10s|\n", s)  // |hello     |  (左对齐，总宽度10)
	fmt.Printf("|%.3s|\n", s)   // |hel|       (只显示前3个字符)
	fmt.Printf("|%10.3s|\n", s) // |       hel| (总宽度10，显示前3个字符)

	// 字符串切片
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Printf("%s\n", names) // [Alice Bob Charlie]

	// 使用 %q 给字符串加上引号
	fmt.Printf("%q\n", s) // "hello"
}
