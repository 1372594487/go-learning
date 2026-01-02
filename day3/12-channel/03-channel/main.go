/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2026-01-02 23:45:50
 * @LastEditors: '1372594487
 * @LastEditTime: 2026-01-03 01:37:56
 * @Description:
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channel是Go语言中各个并发执行体间的通信机制，是类型相关的管道，
// 用于在goroutine之间传递数据和同步执行

// 不要通过共享内存莱通信，二英通过通信来共享内存。
// 创建channel
// ch1 := make(chan int)      // 创建一个传递int类型数据的       无缓冲channel
// ch2 := make(chan int, 3) // 创建一个传递int类型数据的容量为3的 有缓冲channel

// 操作
// ch1 <- 42        // 发送数据42到ch1通道
// value := <-ch1   // 从ch1通道中接收数据并赋值给value变量
// close(ch1)       // 关闭ch1通道，表示不再发送数据

//特殊用法
// value,ok := <-ch1 // 从ch1通道中接收数据，如果ch1通道关闭了，ok为false，否则为true

// 缓冲区区别：
// 无缓冲的channel，同步通信，发送和接受必须同时准备好
// 有缓冲的channel，异步通信，缓冲区满时发送阻塞，空时接受阻塞

type Order struct {
	ID        int
	UserID    string
	Amount    float64
	Status    string
	CreatedAt time.Time
}

func orderProduct(orderChan chan<- Order, number int) {
	for i := 0; i < number; i++ {
		order := Order{
			ID:        i,
			UserID:    fmt.Sprintf("user_%d", rand.Intn(100)),
			Amount:    rand.Float64() * 1000,
			Status:    "pending",
			CreatedAt: time.Now(),
		}
		orderChan <- order
		fmt.Printf("生成订单:ID=%d, 用户ID=%s, 金额=%.2f\n", order.ID, order.UserID, order.Amount)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	close(orderChan)
}

func orderProcessor(orderChan <-chan Order, resultChan chan<- Order) {
	for order := range orderChan {
		fmt.Printf("处理订单:ID=%d, 用户ID=%s, 金额=%.2f\n", order.ID, order.UserID, order.Amount)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		order.Status = "completed"
		resultChan <- order
	}
	close(resultChan)
}

func orderResultCollector(resultChan <-chan Order, done chan bool) {
	for order := range resultChan {
		fmt.Printf("订单处理结果:ID=%d, 用户ID=%s, 金额=%.2f, 状态=%s\n", order.ID, order.UserID, order.Amount, order.Status)
	}
	done <- true
}

// main 函数是程序的入口点
// 当程序启动时，会首先执行此函数
func main() {
	// // 模拟生产者、消费者场景
	// // 生产者：生成订单并发送到订单通道
	// // 消费者：从订单通道接收订单，处理后发送到结果通道
	// // 结果收集器：从结果通道接收处理结果并输出
	// rand.Seed(time.Now().UnixNano())
	// orderChan := make(chan Order, 5)
	// resultChan := make(chan Order, 5)
	// done := make(chan bool)

	// go orderProduct(orderChan, 20)

	// for i := 0; i < 3; i++ {
	// 	go orderProcessor(orderChan, resultChan)
	// }

	// go orderResultCollector(resultChan, done)

	// <-done
	// fmt.Println("所有订单处理完成")

	// 多路复用
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	ch1 <- "来自通}道1的数据"
	// }()

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	ch2 <- "来自通道2的数据"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg := <-ch1:
	// 		fmt.Println(msg)
	// 	case msg := <-ch2:
	// 		fmt.Println(msg)
	// 	case <-time.After(time.Second * 3):
	// 		fmt.Println("超时了")
	// 		return
	// 	}
	// }

	// 定时器
	// ticker := time.NewTicker(time.Millisecond * 500)
	// done := make(chan bool)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case t := <-ticker.C:
	// 			fmt.Println("当前时间：", t.Format("2006-01-02 15:04:02"))
	// 		}
	// 	}
	// }()

	// time.Sleep(time.Second * 2)
	// ticker.Stop()
	// done <- true
	// fmt.Println("Ticker已停止")

	// 工作池
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go worker(jobs, results)
	}
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for value := range results {
		fmt.Println("result:", value)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker:", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
