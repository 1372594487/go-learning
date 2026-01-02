/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 16:52:26
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-02 17:19:51
 * @FilePath: \go-learning\day2\08-OOP\03-interface\interface.go
 * @Description: 1、interface接口实现以及多态现象
 *				 2、万能数据类型 interface{} 以及类型断言
 */
package main

import "fmt"

// interface
type Animal interface {
	Speak() string
	GetColor() string
	GetType() string
}
type Dog struct {
	Name  string
	Color string
}

func (d Dog) Speak() string {
	return "汪汪汪"
}

func (d Dog) GetColor() string {
	return d.Color
}
func (d Dog) GetType() string {
	return "狗"
}

type Cat struct {
	Name  string
	Color string
}

func (c Cat) Speak() string {
	return "喵喵喵"
}

func (c Cat) GetColor() string {
	return c.Color
}
func (c Cat) GetType() string {
	return "猫"
}

func animalType(a Animal) string {
	return a.GetType()
}

/* ----------------------------------------------------------------------- */
//万能数据类型 interface{}
func myFunc(arg interface{}) {
	// interface{} 类型断言机制
	switch v := arg.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown")
		fmt.Println("arg =>", arg)
	}
	//或者
	value, ok := arg.(int)
	if ok {
		fmt.Println("int", value, ok)
	} else {
		fmt.Println("unknown")
		fmt.Println("arg =>", arg)
	}
	// 只需要断言而不需要值，可以使用 _ 代替
	// _, ok := arg.(int)
	// if ok {
	// 	fmt.Println("int", ok)
	// } else {
	// 	fmt.Println("unknown")
	// 	fmt.Println("arg =>", arg)
	// }

}
func main() {
	var animal Animal
	animal = &Cat{"小花", "黑色"}
	fmt.Println(animal.Speak())     // 喵喵喵 多态的现象
	fmt.Println(animal.GetType())   // 猫
	fmt.Println(animalType(animal)) // 猫
	animal = &Dog{"小黑", "白色"}
	fmt.Println(animal.Speak())     // 汪汪汪
	fmt.Println(animal.GetType())   // 狗
	fmt.Println(animalType(animal)) // 狗
	fmt.Println("-------------------------")
	myFunc(10)
	myFunc("hello")

}
