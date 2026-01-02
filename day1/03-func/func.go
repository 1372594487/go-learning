/*
  - @Author: zywOo 1372594487@qq.com
  - @Date: 2025-07-01 14:14:54

* @LastEditors: 1372594487 1372594487@qq.com
* @LastEditTime: 2025-12-21 03:51:00
  - @FilePath: \go-learning\day1\func\func.go
  - @Description:函数多返回值写法
    方法定义：
    func (接受者 接受者类型) 方法名(参数列表) (返回值列表) {
    // 方法体
    }
    接受者类型：
    1、值接收者：操作结构体的副本，对副本的修改不会影响原始结构体
    2、指针接收者：操作结构体的指针，对结构体的修改会影响原始结构体
    应用场景：在金融应用中，账户对象需要封装各种操作（存款、取款、查询余额），这些操作都需要修改账户对象的状态，因此应该使用指针接收者
    *
*/
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// 多个返回值的函数,返回值类型用括号括起来
func swap(a, b string) (string, string) {
	return b, a
}

//匿名返回值 ,返回值类型可以省略

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 函数的参数类型可以省略,编译器会自动推断
func add2(a, b int) int {
	return a + b
}

// 函数的参数可以命名,返回值也可以命名
func split2(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 函数的参数可以命名,返回值也可以命名,返回值可以省略
func split3(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

type BankAccount struct {
	AccountNumber string
	AccountHolder string
	balance       float64
	IsActive      bool
}

func (acc BankAccount) GetAccountInfo() string { // 值接收者
	status := "Active"
	if !acc.IsActive {
		status = "Inactive"
	}
	return fmt.Sprintf("Account Number: %s, Account Holder: %s, Balance: %.2f, Status: %s", acc.AccountNumber, acc.AccountHolder, acc.balance, status)
}

func (acc *BankAccount) Deposit(amount float64) error { // 指针接收者
	if !acc.IsActive {
		fmt.Println("Account is inactive, cannot deposit")
		return fmt.Errorf("")
	}
	if amount <= 0 {
		fmt.Println("Invalid amount, cannot deposit")
		return fmt.Errorf("Invalid amount cannot de")
	}
	acc.balance += amount
	return nil
}

func (acc *BankAccount) Withdraw(amount float64) error { // 指针接收者
	if !acc.IsActive {
		return fmt.Errorf("Account is inactive, cannot withdraw")
	}
	if amount <= 0 {
		fmt.Println("Invalid amount, cannot withdraw")
		return fmt.Errorf("Invalid amount, cannot withdraw")
	}
	acc.balance -= amount
	return nil
}

func (acc *BankAccount) Freeze() { // 指针接收者
	acc.IsActive = false
}

func (acc *BankAccount) Unfreeze() { // 指针接收者
	acc.IsActive = true
}

func main() {
	// fmt.Println(add(42, 13))
	// fmt.Println(swap("hello", "world"))
	// fmt.Println(split(17))
	// fmt.Println(add2(42, 13))
	// fmt.Println(split2(17))
	// fmt.Println(split3(17))

	account := BankAccount{
		AccountNumber: "1234567890",
		AccountHolder: "John Doe",
		balance:       1000.00,
		IsActive:      true,
	}
	fmt.Println(account.GetAccountInfo())
	err := account.Deposit(500.0)
	if err != nil {
		fmt.Println(err)
	}
	// 即使你定义的是值类型变量，仍然可以调用需要指针接收者的方法，Go会自动处理这种转换。
	fmt.Println("account after deposit:", account.GetAccountInfo())
}
