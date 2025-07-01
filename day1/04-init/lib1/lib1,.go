package lib1

import "fmt"

// 定义一个函数Lib1_Fn
func Lib1_Fn() {
	// 打印字符串"Lib1_Fn()"
	fmt.Println("Lib1_Fn()")
}
// 初始化函数
func init() {
	fmt.Println("lib1 init()")
}
