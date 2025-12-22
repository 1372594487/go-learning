/*
* @Author: 1372594487 1372594487@qq.com
* @Date: 2025-12-23 01:16:44
  - @LastEditors: '1372594487
  - @LastEditTime: 2025-12-23 01:51:55

* @Description: 接口与结构体组合应用

结构体代替类：Go使用结构体组织数据，方法通过接收者定义
组合优于继承：通过组合结构体实现代码复用和行为扩展，避免传统继承带来的复杂性
方法接收者：值接收者不修改原对象，指针接受者可修改原对象
接口实现多态：不同类型实现相同接口，提高代码灵活性和可扩展性
方法提升：通过组合，内嵌结构体的方法可以直接在外层结构体上调用（自动提升）
应用场景：
通知系统：不同通知方式（邮件、短信）实现统一接口，方便扩展和维护
订单处理：订单服务通过接口调用不同通知方式，实现解耦和灵活配置
*/
package main

import "fmt"

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct {
	SMTPHost string
	Post     int
}
type SmsNotifier struct {
	APIKey   string
	TmplCode string
}

type BroadCastNotifier struct {
	Notifiers []Notifier
}

type OrderService struct {
	Notifier Notifier
}

func (e EmailNotifier) Notify(message string) error {
	fmt.Printf("发送邮件通知:%s\n", message)
	return nil
}

func (s SmsNotifier) Notify(message string) error {
	fmt.Printf("发送短信通知%s\n", message)
	return nil
}

func (b BroadCastNotifier) Notify(message string) error {
	fmt.Printf("发送广播通知:%s\n")
	return nil
}

func (o *OrderService) SetNotifier(n Notifier) {
	o.Notifier = n
}

func (o *OrderService) CreateOrder(product string, quantity int) error {
	fmt.Printf("创建订单:%s x %d\n", product, quantity)
	err := o.Notifier.Notify("订单创建成功")
	if err != nil {
		fmt.Println("通知失败：", err)
		return fmt.Errorf("通知失败：%v", err)
	}
	return nil
}
func notifyDemo() {
	orderService := &OrderService{}
	emailNotifier := EmailNotifier{SMTPHost: "smtp.example.com", Post: 587}
	orderService.SetNotifier(emailNotifier)
	orderService.CreateOrder("iPhone", 10)

	smsNotifier := SmsNotifier{APIKey: "123456", TmplCode: "123456"}
	orderService.SetNotifier(smsNotifier)
	orderService.CreateOrder("iPad", 5)

	broadcastNotifier := BroadCastNotifier{}
	orderService.SetNotifier(broadcastNotifier)
	orderService.CreateOrder("MacBook", 3)
}
