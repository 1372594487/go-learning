/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 16:00:47
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-02 16:05:38
 * @FilePath: \go-learning\day2\08-struct\struct.go
 * @Description: struct定义与使用
 *
 */
package main

import "fmt"

// type关键字定义结构体
type Person struct {
	name string
	age  int
}

func changeName(p Person) {
	p.name = "李四"
}

type Person2 struct {
	name string
	age  int
}

func changeName2(p *Person2, name string, age int) {
	// Go语言中的逻辑或运算符||不能用于赋值操作。Go语言也不支持三元运算符
	if name != "" {
		p.name = name
	}
	if age != 0 {
		p.age = age
	}
}

func main() {
	var p Person
	p.name = "张三"
	p.age = 18
	fmt.Println(p)
	fmt.Printf("p的类型是：%T\n", p)
	changeName(p)
	fmt.Println(p)
	/* ----------------------------------------------------------------------- */
	fmt.Println("-------------------------")
	var p2 Person2
	p2.name = "王五"
	p2.age = 18
	fmt.Println(p2)
	changeName2(&p2, "赵六", 20)
	fmt.Println(p2)
}
