/*
* @Author: 1372594487 1372594487@qq.com
* @Date: 2026-01-05 11:58:08
 * @LastEditors: '1372594487
 * @LastEditTime: 2026-01-05 14:43:56
* @Description: 并发模式（生产者-消费者、扇入/扇、Pipeline）
应用场景：
生产者-消费者适用场景：
1、消息队列系统（订单处理、日志收集）
2、文件上传后的异步处理
3、爬虫系统的URL调度与页面解析


扇入/扇出适用场景：
1、大数据处理的并行计算
2、负载均衡分发任务
3、多源数据聚合（监控数据汇总）

Pipeline模式适用场景：
1、数据处理流水线（ETL过程）
2、图像处理的多阶段操作（如滤镜、裁剪、缩放等）
3、编译器的工作流程（词法分析、语法分析、优化、生成代码）
*/

// 生产者-消费者模式
// 解决生产者和消费者速度不匹配的问题
// 扇入/扇出模式
// 扇出：一个channel分发给多个goroutine处理（一产多消）
// 扇入：多个channel合并到一个channel（多产一消）
// Pipeline模式
// 多个阶段，每个阶段由一个或多个goroutine处理，前一个阶段的输出作为后一个阶段的输入

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 生产者-消费者模型
type Order struct {
	ID        int
	UserID    int
	Amount    float64
	status    string
	CreatedAt time.Time
}

func orderProduct(orderChan chan<- Order, number int) {
	defer close(orderChan)
	for i := 0; i < number; i++ {
		order := Order{
			ID:        i + 1,
			UserID:    rand.Intn(1000) + 1,
			Amount:    rand.Float64() * 1000,
			status:    "pending",
			CreatedAt: time.Now(),
		}
		fmt.Printf("生成订单：ID=%d,金额=￥%.2f\n", order.ID, order.Amount)
		orderChan <- order
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func orderConsumer(orderChan <-chan Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChan {
		// 模拟订单处理
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		order.status = "completed"
		fmt.Printf("处理订单：ID=%d,状态=%s\n", order.ID, order.status)
	}

}

// 扇出
func funOutProcessor(orderChan <-chan Order, orderype int, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChan {
		switch orderype {
		case 1: //计算折扣
			fmt.Printf("订单处理：ID=%d,状态=%s\n", order.ID, order.status)
		case 2: //发送通知
			fmt.Printf("订单处理：ID=%d,状态=%s\n", order.ID, order.status)
		case 3: //记录日志
			fmt.Printf("订单处理：ID=%d,状态=%s\n", order.ID, order.status)
		}
		time.Sleep(time.Millisecond * 100)
	}

}

// Pipeline模式
func validationStage(in <-chan Order) <-chan Order {
	output := make(chan Order, 10)
	go func() {
		defer close(output)
		for order := range in {
			// 第一阶段，订单验证
			if order.Amount > 0 {
				order.status = "validated"
				fmt.Printf("验证通过：订单%d\n", order.ID)
				output <- order
			} else {
				fmt.Printf("验证失败：订单%d金额异常\n", order.ID)
			}
		}
	}()
	return output
}

func paymentStage(in <-chan Order) <-chan Order {
	output := make(chan Order, 10)
	go func() {
		defer close(output)
		for order := range in {
			// 第二阶段，支付处理
			time.Sleep(time.Millisecond * 200) // 模拟支付处理时间
			order.status = "paid"
			fmt.Printf("支付完成：订单%d\n", order.ID)
			output <- order
		}
	}()
	return output
}

func shippingStage(input <-chan Order) <-chan Order {
	output := make(chan Order, 10)
	go func() {
		defer close(output)
		for order := range input {
			time.Sleep(time.Millisecond * 150)
			order.status = "shipped"
			fmt.Printf("发货完成：订单%d\n", order.ID)
			output <- order
		}
	}()
	return output
}
func main() {
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("生产者-消费者模式示例")
	// orderCh := make(chan Order, 5)
	// var wg sync.WaitGroup

	// //启动消费者
	// wg.Add(2)
	// go orderConsumer(orderCh, &wg)
	// go orderConsumer(orderCh, &wg)

	// //启动生产者
	// orderProduct(orderCh, 6)
	// wg.Wait()

	// fmt.Println("扇入/扇出模式示例")
	// // 扇出 一个输入，多个处理
	// fanOutCh := make(chan Order, 10)
	// var fanOutWg sync.WaitGroup

	// // 启动多个处理goroutine
	// fanOutWg.Add(3)
	// for i := 1; i <= 3; i++ {
	// 	go funOutProcessor(fanOutCh, i, &fanOutWg)
	// }

	// go func() {
	// 	for i := 1; i <= 6; i++ {
	// 		fanOutCh <- Order{ID: i, Amount: float64(i) * 100.0, status: "pending"}
	// 	}
	// 	close(fanOutCh)
	// }()
	// fanOutWg.Wait()

	fmt.Println("Pipeline模式示例")
	// 创建初始输入
	pipelineInput := make(chan Order, 10)

	// 构建流水线
	validatedOrders := validationStage(pipelineInput)
	paidOrders := paymentStage(validatedOrders)
	shippedOrders := shippingStage(paidOrders)

	//发送订单到流水线
	go func() {
		for i := 1; i <= 3; i++ {
			pipelineInput <- Order{ID: i, UserID: i * 10, Amount: float64(i) * 100.0, status: "pending"}
		}
		close(pipelineInput)
	}()

	// 等待流水线处理完成
	fmt.Println("所有订单处理完成")
	for order := range shippedOrders {
		fmt.Printf("订单最终状态：ID=%d,状态=%s\n", order.ID, order.status)
	}

}
