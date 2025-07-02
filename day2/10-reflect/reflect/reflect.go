/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 23:22:34
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-03 00:07:36
 * @FilePath: \go-learning\day2\10-reflect\reflect.go
 * @Description: reflect基本用法
 *
 */
package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Field1 int
	Field2 string
}

// (s MyStruct)表示这个方法是绑定到MyStruct类型上的，也就是说，这个方法只能被MyStruct类型的变量调用。s是这个方法的接收者，它是一个MyStruct类型的变量。
func (s MyStruct) Method1(x int) int {
	return s.Field1 + x
}

func main() {
	//1 获取类型信息
	var a float64 = 3.14
	fmt.Println("type:", reflect.TypeOf(a))
	//2 获取值信息
	v := reflect.ValueOf(a)
	fmt.Println("value:", v)
	//3 修改值 要修改反射对象的值，需要先确保它是“可寻址”的（即通过指针获取反射对象），然后使用Set()方法。
	p := reflect.ValueOf(&a) //获取a的地址
	v2 := p.Elem()           //获取指针指向的值
	v2.SetFloat(1.23)
	fmt.Println("value:", v2)
	//链式调用
	p2 := reflect.ValueOf(&a).Elem()
	p2.SetFloat(1.24)
	fmt.Println("value:", p2)
	//4 获取结构体信息
	s := MyStruct{10, "hello"}
	v3 := reflect.ValueOf(s)
	fmt.Println("Field1:", v3.Field(0))
	fmt.Println("Field2:", v3.Field(1))

	//5 调用方法
	// 创建一个结构体实例
	s2 := MyStruct{10, "world"}
	// 获取结构体的反射值
	v4 := reflect.ValueOf(s2) //{10 world}
	// 获取结构体的第一个方法
	mv := v4.Method(0) //0x209640
	fmt.Println("Method:", mv, "value:", v4)
	// 调用方法，并传入参数
	// []reflect.Value：这是一个reflect.Value类型的切片，表示要传递给方法的参数列表。每个reflect.Value元素代表一个参数的值。
	res := mv.Call([]reflect.Value{reflect.ValueOf(5)})
	// 如果你直接写res := mv.Call([]reflect.Value{5})，那么5将被视为一个整数，而不是reflect.Value类型的值。
	// 这将导致编译错误，因为Call方法期望的参数类型是[]reflect.Value，而不是直接的原始值。
	// 所以，你需要使用reflect.ValueOf函数将原始值转换为reflect.Value类型的值，然后再传递给Call方法。
	// 打印结果
	fmt.Println("Result:", res[0].Int())
	// 以上只是reflect包的一些基本用法，它还提供了更多高级功能，如类型断言、接口转换等。
	// 在实际使用中，应谨慎使用反射，因为它可能导致代码难以理解和维护，且可能带来性能开销。

	//

}
