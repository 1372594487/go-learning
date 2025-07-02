/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 12:48:45
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-02 13:22:30
 * @FilePath: \go-learning\day2\07-map\map1\map.go
 * @Description: map的声明以及操作
 *
 */
package main

import "fmt"

func main() {
	// map声明方式

	// 1 var 变量名 map[keyType]valueType
	// var m1 map[string]int
	// m1["no1"] = 100 //panic: assignment to entry in nil map
	// 使用前要给map分配数据空间

	// 2 var 变量名 map[keyType]valueType = make(map[keyType]valueType, capacity)
	var m2 map[string]int = make(map[string]int, 10)
	m2["no1"] = 100
	m2["no2"] = 200
	m2["no3"] = 300

	// 3 变量名 := map[keyType]valueType{key1: value1, key2: value2, ...}
	m3 := map[string]int{"no1": 100, "no2": 200, "no3": 300}
	fmt.Println(m2, m3) // map[no1:100 no2:200 no3:300] map[no1:100 no2:200 no3:300]

	// map的值可以是任意类型，包括map类型
	m4 := map[string]map[string]string{
		"no1": {
			"name": "张三",
			"sex":  "男",
		},
		"no2": {
			"name": "李四",
			"sex":  "女",
		},
	}
	fmt.Println(m4)
	/* ----------------------------------------------------------------------- */
	fmt.Println("-------------------------")

	// map的操作
	// 1 判断某个键是否存在
	// 从map中获取键为"no1"的值
	value, ok := m3["no1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no1不存在")
	}
	// 100

	// 2 遍历map
	for k, v := range m3 {
		fmt.Println(k, v)
	}
	// no1 100
	// no2 200
	// no3 300

	// 3 删除某个键值对
	delete(m3, "no1")
	fmt.Println(m3)
	// map[no2:200 no3:300]

	// 4 获取map的长度
	fmt.Println(len(m3))
	// 2

	// 5 map的值是引用类型，所以直接修改map的值，会影响到所有引用该map的变量
	m3["no2"] = 2000
	fmt.Println(m3)
	// map[no2:2000 no3:300]

}
