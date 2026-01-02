/*
  - @Author: zywOo 1372594487@qq.com

  - @Date: 2025-07-03 15:02:48

  - @LastEditors: '1372594487

  - @LastEditTime: 2026-01-02 02:31:40

  - @FilePath: \go-learning\day3\11-goroutine\goroutine.go

  - @Description:goroutine

    // 并发与并行概念：

// 并发：多个任务在同一个时间段内交替执行（看起来同时）
// 并行：多个任务在同一个时间点同时执行（真正意义上的同时）

CSP理论（Communicating Sequential Processes）
核心思想：不要通过共享内存来通信，而是通过通信来共享内存
组成要素：Process(进程/协程) + Channel(通道)

应用场景
Goroutine适用场景：
高并发Web服务器：每个请求一个Goroutine
数据处理流水线：多个处理阶段并行执行
实时消息推送：Websocket连接管理
定时任务调度：后台定时执行任务
并发爬虫：同时抓取多个网页

	*
*/
package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
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

func fetchData(url string, wg *sync.WaitGroup) {

	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)

	fmt.Println("Fetching data from:", url, "delay", delay)

	if wg != nil {
		defer wg.Done()
	}

}

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "done")
			return
		case <-time.After(time.Second * 1):
			fmt.Println("worker", id, "timeout")
			return
		default:
			fmt.Println("worker", id, "is working")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func limitedWorkerPool(tasks []string, max int) {
	start := time.Now()
	semaphore := make(chan struct{}, max)
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(i int, task string) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			var j = i + 1
			fmt.Println("task", j, "start")
			time.Sleep(time.Second * 2)
			fmt.Println("task", j, "done")
		}(i, task)
	}
	wg.Wait()
	fmt.Println("Total time taken :", time.Since(start))
}

// 运行结果示例：
// 假设 tasks 有 5 个元素，max 为 2：
// 任务 0 和 1 立即开始。
// 任务 2、3、4 会阻塞在 semaphore <- struct{}{} 这一行。
// 约 2 秒后，任务 0 和 1 完成，释放槽位。
// 任务 2 和 3 获得槽位并开始。
// 再过约 2 秒，任务 2 和 3 完成，任务 4 开始。
// 最后任务 4 完成。
// 总耗时约为 6 秒（2秒 * 3批次），而不是串行的 10 秒或无限制并发的 2 秒。

func showGoroutineInfo() {
	fmt.Printf("当前Goroutine数量：%d\n", runtime.NumGoroutine())
	fmt.Printf("当前CPU核心数:%d\n", runtime.NumCPU())

	// 设置最大并行数
	runtime.GOMAXPROCS(4)
}

// 主 goroutine
func main() {
	// go newTask() // 启动一个 goroutine 执行 newTask 函数

	start := time.Now()
	// var wg sync.WaitGroup
	urlSlice := []string{
		"http://www.baidu.com", "http://www.google.com", "http://www.github.com", "http://www.zhihu.com", "http://www.weibo.com"}
	for _, v := range urlSlice {
		// wg.Add(1)
		fetchData(v, nil)
	}
	fmt.Println("Total time taken :", time.Since(start))
	start = time.Now()
	fmt.Println("==============================================")
	var wg sync.WaitGroup
	for _, v := range urlSlice {
		wg.Add(1)
		go fetchData(v, &wg)
	}
	wg.Wait()
	fmt.Println("Total time taken :", time.Since(start))

	fmt.Println("==============================================")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(time.Second * 5)

	limitedWorkerPool(urlSlice, 3)

	showGoroutineInfo()
}
