/*
  - @Author: zywOo 1372594487@qq.com
  - @Date: 2025-07-03 15:37:03

* @LastEditors: zywOo 1372594487@qq.com
* @LastEditTime: 2025-07-03 17:24:31
* @FilePath: \go-learning\day3\12-channel\channel.go
  - @Description:1 channel基本定义与使用
    2 channel的关闭特点
    *
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// channel定义方式
	// 未初始化的通道是 nil，向 nil 通道发送数据会导致死锁。	// 从 nil 通道接收数据也会导致死锁。
	// var ch1 chan int
	ch2 := make(chan int)
	ch3 := make(chan int, 3) //带缓冲区的通道
	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行")
		// ch1 <- 666 //向ch1中写入数据
		ch2 <- 777
		//当channel被写满时，会阻塞
		for i := 0; i < 3; i++ {
			ch3 <- i
			fmt.Println("向ch3中写入数据", i, "len(i):", len(ch3), "cap(i):", cap(ch3))
		}
	}()
	//接收数据并赋值
	// data1 := <-ch1
	data2 := <-ch2
	fmt.Println("从ch2中读取数据", data2)
	time.Sleep(2 * time.Second)
	// channel为空，从里边取数据会导致阻塞
	for i := 0; i < 3; i++ {
		data3 := <-ch3
		fmt.Println("从ch3中读取数据", data3, "len(i):", len(ch3), "cap(i):", cap(ch3))
	}
	fmt.Println("main goroutine 结束")
	/* ----------------------------------------------------------------------- */
	// channel关闭
	// 关闭channel后，无法向channel中写入数据，但可以读取数据
	// 关闭channel后，从channel中读取数据时，如果channel中为空，则返回对应类型的零值
	// 关闭channel后，从channel中读取数据时，如果channel中不为空，则可以读取到数据
	// 关闭channel后，多次关闭会导致panic
	// 关闭channel后，向channel中写入数据会导致panic
	// 对于 nil channel，无论执行读操作还是写操作都会被阻塞
	c := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}
		//关闭channel 发送完关闭，否则会造成死锁
		close(c)
	}()
	for {
		//如果channel被关闭，则ok为false 如果channel未被关闭，则ok为true
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
		// 用range迭代不断操作channel
		// for data := range c {
		// 	fmt.Println(data)
		// }
	}
	fmt.Println("main goroutine 结束")
}
