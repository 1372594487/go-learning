/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-03 00:18:22
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-03 00:30:29
 * @FilePath: \go-learning\day2\10-reflect\struct_tag\struct_tag.go
 * @Description:结构体标签以及在json中的应用
 *
 */
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 反射解析结构体标签
type student struct {
	Name string `info:"name" doc:"我的名字"`
	Age  int    `info:"sex"`
}

// 结构体标签在json中的应用
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func findTag(stru interface{}) {
	t := reflect.TypeOf(stru) //reflect.Type
	// 使用NumField()方法获取结构体中的字段数量，然后通过循环遍历每个字段
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("tagInfo:", tagInfo, "tagDoc:", tagDoc)
	}
}

func main() {
	stu := student{
		Name: "张三",
		Age:  18,
	}
	findTag(stu)
	/* ----------------------------------------------------------------------- */
	movie := Movie{
		Title:  "Avengers",
		Year:   2019,
		Price:  10,
		Actors: []string{"Robert Downey Jr.", "Chris Evans", "Mark Ruffalo"},
	}
	//编码过程 结构体 -> json
	jsonStr, err := json.Marshal(movie) //返回json字符串
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("jsonStr: %s\n", jsonStr)
	// 解码过程 json -> 结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Printf("%#v\n", myMovie)
}
