/*
- @Author: zywOo 1372594487@qq.com

- @Date: 2025-07-02 12:48:45

  - @LastEditors: 1372594487 1372594487@qq.com

  - @LastEditTime: 2025-12-21 02:40:22

- @FilePath: \go-learning\day2\07-map\map1\map.go

  - @Description:
    并发安全问题Go的map是线程安全的，在并发读写时不是线程安全，需要配合sync包中的锁；或者使用sync.Map
    应用场景
    1.缓存系统：将数据缓存到内存中，提高数据访问速度
    2.计数器：统计词频、某些事件的发生次数
    3.数据字典（数据库查询结果索引）：存储键值对数据，快速查找
    4.数据分组：将数据按照某种规则分组，便于统计和分析
    5.JSON数据解析：处理JSON数据时，将JSON数据解析为map类型，方便操作

    *
*/
package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {

	//映射是内置数据结构，用于存储键值对的无序集合，映射是引用类型，传递开销小
	//映射声明与初始化

	// var m1 map[string]int // 声明nil映射，键为string类型，值为int类型
	// var 变量名 map[keyType]valueType
	// var m1 map[string]int
	// m1["no1"] = 100 //panic: assignment to entry in nil map
	// 使用前要给map分配数据空间

	m2 := make(map[string]int)                               // 使用make声明并初始化映射
	m3 := map[string]int{"no1": 100, "no2": 200, "no3": 300} // 字面量创建
	// map的值可以是任意类型，包括map类型
	// m4 := map[string]map[string]string{
	// 	"no1": {
	// 		"name": "张三",
	// 		"sex":  "男",
	// 	},
	// 	"no2": {
	// 		"name": "李四",
	// 		"sex":  "女",
	// 	},
	// }

	//操作
	m2["apple"] = 5             //插入或更新
	value := m2["apple"]        //查找
	delete(m2, "apple")         //删除
	value, exist := m2["apple"] //检查键是否存在
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("apple不存在")
	}
	// 从map中获取键为"no1"的值
	value, ok := m3["no1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no1不存在")
	}
	// 100

	//遍历map
	for k, v := range m3 {
		fmt.Println(k, v)
	}
	// no1 100
	// no2 200
	// no3 300

	// 获取map的长度
	fmt.Println(len(m3))
	// 2

	//map的值是引用类型，所以直接修改map的值，会影响到所有引用该map的变量
	m3["no2"] = 2000
	fmt.Println(m3)
	// map[no2:2000 no3:300]

	//作为集合使用
	uniqueIDs := map[int]bool{}
	ids := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	for _, id := range ids {
		if _, ok := uniqueIDs[id]; !ok {
			uniqueIDs[id] = true
		}
	}
	fmt.Printf("唯一ID数量： %d\n", len(uniqueIDs))

	//键值反转
	originMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	reverseMap := make(map[int]string)
	for k, v := range originMap {
		reverseMap[v] = k
	}
	// map 是无序的键值对
	ageMap := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	keys := make([]string, 0, len(ageMap))
	for key := range ageMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("Name: %s, Age: %d\n", key, ageMap[key])
	}

	//sync.Map
	//sync.Map 是 Go 语言标准库中的一个并发安全的映射类型，它提供了一些方法来安全地存储和读取键值对，而不需要使用互斥锁来保护并发访问。
	var syncMap sync.Map
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")

	fmt.Println(syncMap.Load("key1")) // 输出: value1
	syncMap.Delete("key1")

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})

}
