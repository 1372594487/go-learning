/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-01 14:54:40
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-01 16:04:40
 * @FilePath: \go-learning\day1\04-init\main.go
 * @Description:匿名以及别名导包方式
 *
 */
package main

import (
	"github.com/1372594487/go-learning/day1/04-init/lib1"
	"github.com/1372594487/go-learning/day1/04-init/lib2"
	//匿名导包方式 无法使用包中的函数但会会执行包中的init函数
	// _ "github.com/1372594487/go-learning/day1/04-init/lib1"
	//别名导包方式通过别名调用包中的函数
	// lib2 "github.com/1372594487/go-learning/day1/04-init/lib2"
	// 包中存在init函数，在包被引用时会自动执行init函数
)

// 主函数
func main() {
	// 调用lib1包中的Lib1_Fn函数
	lib1.Lib1_Fn()
	// 调用lib2包中的Lib2_Fn函数
	lib2.Lib2_Fn()
}
