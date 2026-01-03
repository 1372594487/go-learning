/*
* @Author: 1372594487 1372594487@qq.com
* @Date: 2026-01-04 00:37:55
  - @LastEditors: '1372594487
  - @LastEditTime: 2026-01-04 02:34:33

* @Description: 同步原语（Mutex RWMutex WaitGroup etc.）
适用场景
sync.Mutex适用场景：
银行账户余额更新（防止超额取款）
电商库存扣减（防止超卖）
计数器累加操作

sync.WaitGroup适用场景：
批量文件下载，等待所有下载完成
并行数据处理，等待所有处理结果
服务启动时等待多个初始化任务完成

sync.RWMutex适用场景：
配置信息读取（频繁读取，偶尔更新）
缓存系统（大量查询，少量更新）
热点数据访问（如商品详情页）

使用原则：
尽量减少锁的粒度，避免不必要的锁竞争
使用defer确保锁的释放，防止死锁
避免在持有锁时执行耗时操作
读写分离场景优先使用sync.RWMutex提高并发性能
简单场景优先使用原子锁sync/atomic提高性能
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 并发编程中，多个goroutine同时访问共享资源时，可能会导致数据竞争和不一致的问题。
// 为了解决这些问题，Go语言提供了一些同步原语，如互斥锁（sync.Mutex）、读写锁（sync.RWMutex）和等待组（sync.WaitGroup）等。

// 互斥锁（sync.Mutex）
// 保证同一时间只有一个2goroutine能访问共享资源
// 两个方法：Lock() 和 Unlock()
// 使用后必须释放，否则会导致死锁

// 读写锁（sync.RWMutex）
// 允许多个读操作或一个写操作
// 读锁方法：RLock()、RUnlock() （共享锁）
// 写锁方法：Lock()、Unlock()   （互斥锁）

// 等待组（sync.WaitGroup）
// 等待一组goroutine完成
// 三个方法：Add()、Done()、Wait()
// Add(n int)：设置等待的goroutine数量
// Done()：表示一个goroutine完成
// Wait()：阻塞直到所有goroutine完成
// 用于协调多个goroutine的执行程序

type Inventory struct {
	stock   int
	rwMutex sync.RWMutex // 读写锁
}

func (inv *Inventory) GetStock() int {
	inv.rwMutex.RLock()
	defer inv.rwMutex.RUnlock()
	return inv.stock
}

func (inv *Inventory) deductStock(quantity int) bool {
	inv.rwMutex.Lock()
	defer inv.rwMutex.Unlock()
	if inv.stock < quantity {
		fmt.Println("库存不足")
		return false
	}
	time.Sleep(time.Millisecond * 100)
	inv.stock -= quantity
	return true
}

type Counter struct {
	value int64
}

// 原子锁
func (v *Counter) Increment() {
	time.Sleep(time.Millisecond * 15)
	atomic.AddInt64(&v.value, 1)
}

func (v *Counter) Decrement() {
	time.Sleep(time.Millisecond * 20)
	atomic.AddInt64(&v.value, -1)
}

func (v *Counter) GetValue() int64 {
	return atomic.LoadInt64(&v.value)
}
func main() {
	// // 模拟电商库存扣减
	// Inventory := &Inventory{stock: 100}
	// // 初始化库存为100
	// var wg sync.WaitGroup

	// // 模拟多个goroutine同时读取库存
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		Inventory.deductStock(1)
	// 	}(i)
	// }

	// for i := 0; i < 20; i++ {
	// 	time.Sleep(time.Millisecond * 100)
	// 	fmt.Printf("Goroutine %d 读取库存: %d\n", i, Inventory.GetStock())

	// }
	// wg.Wait()
	// fmt.Println("最终库存:", Inventory.GetStock())

	// 点赞
	var counter Counter
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Decrement()
		}()
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Printf("Goroutine %d 读取点赞数: %d\n", i, counter.GetValue())
	}
	wg.Wait()
	fmt.Println("最终点赞数:", counter.GetValue())

}
