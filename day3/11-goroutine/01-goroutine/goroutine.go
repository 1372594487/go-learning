/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-03 15:02:48
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-03 15:14:53
 * @FilePath: \go-learning\day3\11-goroutine\goroutine.go
 * @Description:goroutine
 *
 */
package main

import (
	"fmt"
	"time"
)

// goroutine 是由 Go 运行时系统管理的轻量级线程
// 定义一个函数，用于创建一个新任务
// 子goroutine
func newTask() {
	// 初始化一个变量i为0
	i := 0
	// 无限循环
	for {
		// 每次循环i加1
		i++
		// 打印出当前i的值
		fmt.Printf("newTask: %d\n", i)
		// 暂停1秒钟
		time.Sleep(1 * time.Second)
	}
}

// 主 goroutine
func main() {
	go newTask() // 启动一个 goroutine 执行 newTask 函数
}
