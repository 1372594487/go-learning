/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-03 17:35:53
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-03 17:41:52
 * @FilePath: \go-learning\day3\12-channel\02-channel\channel.go
 * @Description:chennel 与 select
 *
 */
package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		//select 可以同时监听多个channel
		select {
		case c <- x:
			x, y = y, x+y // x = y,y = x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacii(c, quit)
}
