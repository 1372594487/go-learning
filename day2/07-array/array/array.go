/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 11:02:19
 * @LastEditors: 1372594487 1372594487@qq.com
 * @LastEditTime: 2025-12-17 20:34:11
 * @FilePath: \go-learning\day2\07-array\array.go
 * @Description:array声明，赋值，遍历
 应用场景
  1、存储固定大小的数据集合（一周温度数据）
  2、内存敏感场景，数组在栈上分配
  3、加密/编码等需要精确控制内存布局的操作
 *
 */
package main

import "fmt"

// 数组
// 数组是值类型，意味着数组被复制给一个新变量或传递给函数时，会创建整个数组的副本
// 数组是定长的
// 数组是连续的内存空间

//数组声明方式
var arr1 [5]int //声明但不初始化，默认值为0
var arr2 = [3]int{1, 3, 5}	//声明并初始化
var arr3 = [...]int{2, 4, 6, 8, 10} //声明并初始化，...表示根据初始化的值自动计算数组长度

// 数组是值类型，赋值和传参会复制整个数组，而不是指针
func printArray(arr [5]int) {
	// 值拷贝
	for i, v := range arr {
		fmt.Println("index=>", i, "value=>", v)
	}
}

// 二维数组
var grid = [3][5]int{
	{0, 1, 2, 3, 4},
}

func printGrid(arr [3][5]int) {
	// 值拷贝
	for i, v := range arr {
		fmt.Println("index=>", i, "value=>", v)
	}
}

// 动态数组
var arr5 = []int{1, 3, 5}

// 动态数组遍历
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
	for i, v := range s {
		fmt.Println("index=>", i, "value=>", v)
	}
}

// 举例
func updateArray(arr *[5]int) {
	arr[0] = 99
}

func main() {
	// printArray(arr1)
	// printArray(arr2)
	//cannot use arr2 (variable of type [3]int) as [5]int value in argument to printArray
	printSlice(arr2[:])
	/*
		arr2[:]这样的语法，它实际上是在创建一个基于数组arr2的新切片，这个切片包含了数组arr2的所有元素。
		以下是使用切片而不是直接使用数组的一些原因：
			性能优化：如果数组很大，而你只需要数组的一小部分，传递整个数组可能会浪费内存和CPU时间。切片只包含它所引用的数据的引用，因此它通常比传递整个数组更高效。

			动态大小：切片的长度可以动态调整，而数组的长度是固定的。这使得切片在处理变化的数据集时更加灵活。

			函数参数：在Go语言中，函数参数是通过值传递的。这意味着如果你传递一个数组给函数，实际上会复制整个数组。这可能会导致不必要的性能开销。相反，传递一个切片只会复制切片本身（这是一个很小的结构），而不是底层数组。

			接口兼容性：许多Go标准库函数和用户定义的函数接受切片作为参数，而不是数组。这是因为切片更通用，可以用来表示任意大小的数据集。

	*/

}
