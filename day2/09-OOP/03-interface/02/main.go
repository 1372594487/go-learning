/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-22 19:02:21
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-22 21:02:06
 * @Description:
 应用场景：
 在实际电商系统中，需要接入多种支付方式，如微信支付、支付宝支付、银联支付等。每种支付方式都有相同的核心操作
 （发起支付、查询支付状态、退款等），但具体实现不同。通过定义一个支付接口，可以让不同的支付方式实现该接口，
 从而实现统一的支付处理逻辑，方便扩展和维护。
*/

package main

import (
	"fmt"
	"time"
)

//接口是Go语言中实现多态的重要特性，它定义了一组方法签名，
// 任何类型只要实现了这些方法，就被认为实现了该接口。无需显示声明
// type 接口名 interface {
//     方法1(参数列表) (返回值列表)
// 	方法2(参数列表) (返回值列表)
// 	...
// }

type Payment interface {
	InitiatePayment(amount float64) (string, error)
	RefundPayment(transactionID string, amount float64) (string, error)
	// QueryPaymentStatus(transactionID string) (string, error)
}
type Queryer interface {
	// 假设上边的方法抽理出来一个新的接口
	QueryPaymentStatus(transactionID string) (string, error)
}

// 组合接口
type PaymentService interface {
	Payment
	Queryer
}
type Alipay struct {
	AppID      string
	AppSecret  string
	MerchantID string
}

type WeChatPay struct {
	AppID      string
	AppSecret  string
	MerchantID string
}

// 实现Payment接口的WeChatPay类型的方法
func (a Alipay) InitiatePayment(amount float64) (string, error) {
	return fmt.Sprintf("Alipayd-%d", time.Now().Unix()), nil
}

func (a Alipay) RefundPayment(transactionID string, amount float64) (string, error) {
	return fmt.Sprintf("ALIPAY-REFONC-%s", transactionID), nil

}
func (a Alipay) QueryPaymentStatus(transactionID string) (string, error) {
	return "Success", nil
}

// 实现Payment接口的WeChatPay类型的方法
func (w WeChatPay) InitiatePayment(amount float64) (string, error) {
	return fmt.Sprintf("WeChatPayd-%d", time.Now().Unix()), nil
}

func (w WeChatPay) RefundPayment(transactionID string, amount float64) (string, error) {
	return fmt.Sprintf("WECHATPAY-REFONC-%s", transactionID), nil
}
func (w WeChatPay) QueryPaymentStatus(transactionID string) (string, error) {
	return "Success", nil
}
func ProcessPayment(p PaymentService, amount float64) (string, error) {
	fmt.Println("Processing payment...")
	transactionID, err := p.InitiatePayment(amount)
	if err != nil {
		return "", err
	}
	fmt.Printf("Payment initiated with transaction ID:%s\n", transactionID)
	return transactionID, nil
}

func processEmptyInterface(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer value: %d\n", v)
	case string:
		fmt.Printf("String value: %s\n", v)
	case bool:
		fmt.Printf("Boolean value: %t\n", v)
	default:
		fmt.Println("Unknown type")
	}
}
func main() {
	// 空接口
	// 空接口 interface{} 不包含任何方法，
	var anything interface{}
	processEmptyInterface(anything)
	// 可以赋值给任何类型
	anything = 44 //整数
	processEmptyInterface(anything)
	anything = "hello" //字符串
	processEmptyInterface(anything)
	anything = true //布尔值

	// fmt.Println(anything)

	alipay := Alipay{
		AppID:      "your-alipay-app-id",
		AppSecret:  "your-alipay-app-secret",
		MerchantID: "your-alipay-merchant-id",
	}

	wechatPay := WeChatPay{
		AppID:      "your-wechatpay-app-id",
		AppSecret:  "your-wechatpay-app-secret",
		MerchantID: "your-wechatpay-merchant-id",
	}

	//下边代码等价于先定义一个具名结构体
	// type Transaction struct {
	// 	Payment Payment
	// 	Amount  float64
	// 	Name    string
	// }

	// // 然后使用
	// transactions := []Transaction{
	// 	{Payment: alipay, Amount: 100.0, Name: "Alipay"},
	// 	{Payment: wechatPay, Amount: 200.0, Name: "WeChatPay"},
	// }

	// 定义一个包含Payment接口的切片
	transactions := []struct {
		// Payment
		Payment PaymentService // 这里使用组合接口PaymentService
		Amount  float64
		Name    string
	}{
		{Payment: alipay, Amount: 100.0, Name: "Alipay"},
		{Payment: wechatPay, Amount: 200.0, Name: "WeChatPay"},
	}
	for _, transaction := range transactions {
		transactionID, err := ProcessPayment(transaction.Payment, transaction.Amount)
		if err != nil {
			fmt.Println("Error:", err.Error()) // 处理错误
			continue
		}
		fmt.Printf("%s payment processed successfully, transaction ID: %s\n", transaction.Name, transactionID)
	}

}
