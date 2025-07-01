/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-01 13:33:49
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-01 14:27:56
 * @FilePath: \go-learning\day1\const\const.go
 * @Description: 常量声明
 *
 */

package main

import "fmt"

func main() {
	// 常量声明 (只读属性)
	const pi = 3.14
	// 多常量声明
	const (
		a = 1
		b = 2
		c = 3
	)

	// iota 常量生成器 默认值是0，每新增一行 iota 值加1 只能在const中使用
	const (
		d = iota + 1
		z
		e
		f
	)
	fmt.Println(pi, a, b, c, d, e, f)

}
