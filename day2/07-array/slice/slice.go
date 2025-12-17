/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 11:45:46
 * @LastEditors: 1372594487 1372594487@qq.com
 * @LastEditTime: 2025-12-17 20:49:44
 * @FilePath: \go-learning\day2\07-array\slice\slice.go
 * @Description:切片(slice)声明，以及切片容量机制
 适用场景
 * 1. 动态数据集合（如读取未知长度的文件内容）
 * 2. 复用内存，需要操作部分数据，而不需要复制整个数组的情况（如截取日志子集）
 * 3. 实现栈、队列等数据结构
 * 4. 多个函数间传递数据，而不需要复制整个数组（如函数间传递切片）
 *
 */
package main

import "fmt"

func main() {

	//切片
	//切片提供动态大小的灵活视图，是引用类型，包含三个组件：指向底层数组的指针、长度和容量
	//切片的扩容机制：当切片容量不足时，Go会自动扩容，容量小于1024时通常翻倍，超过1024时每次增加25%左右
	// 切片声明方式
	var slice1 []int               //声明切片
	slice2 := []int{1, 2, 3, 4, 5} //字面量声明
	slice3 := make([]int, 3,5)       //make函数声明，长度3，容量5
	arr := [5]int{1, 2, 3, 4, 5}
	slice4 := arr[1:3] //从数组创建，从数组arr的索引1开始，到索引3结束，不包括索引3
	fmt.Printf("%v\n", slice1) //[]
	fmt.Printf("%v\n", slice2) //[1,2,3,4,5]

	slice5 := slice2[1:3] // 从slice2的索引1开始，到索引3结束，不包括索引3
	slice6 := slice2[:3]  // 从slice2的索引0开始，到索引3结束，不包括索引3
	slice7 := slice2[:]   // 从slice2的索引0开始，到索引len(slice2)结束

	fmt.Printf("%v\n", slice5) // [2 3]
	fmt.Printf("%v\n", slice6) // [1 2 3]
	fmt.Printf("%v\n", slice7) // [1 2 3 4 5]

	/* ----------------------------------------------------------------------- */
	fmt.Println("-------------------------")

	slice3 := make([]int, 3)    //声明一个slice并指定长度,容量默认为长度，值为0
	slice4 := make([]int, 3, 5) //声明一个slice并指定长度和容量

	fmt.Printf("%v\n", slice3) //[0,0,0]
	fmt.Printf("%v\n", slice4) //[0,0,0]

	/* ----------------------------------------------------------------------- */
	fmt.Println("-------------------------")

	slice8 := append(slice2, 4) // 将4添加到slice2的末尾，返回一个新的切片
	slice9 := copy(slice2, slice3)  // 将slice3的内容复制到slice2中，返回复制的元素个数
	slice10 := len(slice2)      // 返回slice2的长度
	slice11 := cap(slice2)      // 返回slice2的容量

	fmt.Printf("%v\n", slice8)  // [1 2 3 4 5 4]
	fmt.Printf("%v\n", slice9)  // 3
	fmt.Printf("%v\n", slice10) // 6
	fmt.Printf("%v\n", slice11) // 5
	/* ----------------------------------------------------------------------- */
	fmt.Println("-------------------------")
	// 切片容量机制
	slice12 := []int{1, 2, 3, 4, 5}
	slice13 := slice12[1:3]             // 从slice12的索引1开始，到索引3结束，不包括索引3
	slice13 = append(slice13, 6)        // 将6添加到slice13的末尾，返回一个新的切片 [2 3 6]
	slice13_len := len(slice13)         // 返回slice13的长度
	slice13_cap := cap(slice13)         // 返回slice13的容量
	fmt.Printf("%v\n", slice13)     // [2 3 6]
	fmt.Printf("%v\n", slice13_len) // 3
	fmt.Printf("%v\n", slice13_cap) // 4
	/* ----------------------------------------------------------------------- */
	fmt.Println("-------------------------")
	slice13 = append(slice13, 7, 7, 7)  // 将7添加到slice13的末尾，返回一个新的切片 [2 3 6 7]
	fmt.Printf("%v\n", slice13)     //[2 3 6 7 7 7]
	fmt.Printf("%v\n", slice13_len) // 3
	fmt.Printf("%v\n", slice13_cap) // 4
	slice13_len = len(slice13)          // 返回slice13的长度
	slice13_cap = cap(slice13)          // 返回slice13的容量
	fmt.Printf("%v\n", slice13_len) // 6
	// 当slice13的长度超过slice12的容量时，slice13会自动扩容，扩容后的容量是原来的两倍
	fmt.Printf("%v\n", slice13_cap) // 8

}
