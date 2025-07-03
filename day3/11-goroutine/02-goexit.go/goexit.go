/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-03 15:18:35
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-03 15:25:36
 * @FilePath: \go-learning\day3\goexit.go\goexit.go
 * @Description:goroutine 退出
 *
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 用go创建一个形参为空，返回值为空的函数
	go func() {
		defer fmt.Println("A defer")
		func() {
			defer fmt.Println("B defer")
			// goexit退出当前goroutine
			runtime.Goexit()
			//fmt.Println("B")不会执行
			fmt.Println("B")
		}()
		// fmt.Println("A")也不会执行 因为Goexit退出了当前goroutine
		fmt.Println("A")
	}()
	//go创建带参数的函数
	go func(name string) {
		fmt.Println(name)
	}("hello")

	for {
		time.Sleep(1 * time.Second)
	}
}
