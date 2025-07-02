/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 17:46:12
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-02 18:04:25
 * @FilePath: \go-learning\day2\10-reflect\pair_test2\pair_test2.go
 * @Description:pair
 *
 */
package main

import (
	"fmt"
	"reflect"
)

// pair
// pair 是一个结构体，包含两个字段：Type 和 Value。Type 字段表示 pair 的类型，Value 字段表示 pair 的值。
type pair struct {
	Type  reflect.Type
	Value reflect.Value
}

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (b Book) ReadBook() {
	fmt.Println("read book")
}

func (b Book) WriteBook() {
	fmt.Println("write book")
}

func main() {
	//b:pair<type:Book, value:&Book{}(0xc0000a6000)>
	b := &Book{}
	//r:pair<type:, value:>
	var r Reader
	//r:pair<type:Book, value:&Book{}(0xc0000a6000)>
	r = b
	r.ReadBook()
	var w Writer
	//r:pair<type:Book, value:&Book{}(0xc0000a6000)>
	w = r.(Writer) // 使用类型断言 r.(Writer) 将 Reader 类型的 r 转换为 Writer 类型，并赋值给 w。类型断言用于将一个接口类型转换为另一个接口类型或具体类型。
	w.WriteBook()
	fmt.Println(w)
}
