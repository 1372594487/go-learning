/*
* @Author: 1372594487 1372594487@qq.com
* @Date: 2025-12-30 14:11:19
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-30 18:15:39
 * @Description:泛型:编写类型安全的通用代码
关键概念：
类型参数：函数或类型定义中的占位类型
类型约束：限制类型参数的范围
类型推断：编译器自动推断类型参数
可比较性：类型是否支持 == 和 != 操作

核心语法：[T 类型约束]、any、comparable、自定义约束
使用原则：真正需要代码复用时使用，不要过度设计

应用场景
1、通用数据结构
栈、队列、链表等容器
集合操作
2、工具函数
比较、排序、查找、过滤等
数学运算
3、数据处理
映射、过滤、转换、聚合等
数据转换
4、避免interface{}和类型断言
类型安全(优势)
更好的性能（优势）
*/

package main

// 定义：编写适用于多种类型的代码，避免重复
// 核心：类型参数+类型约束
// 像写函数的”模版“，可以适配不同类型

// 泛型函数
// func 函数名[T 类型约束](参数列表 T) T {
// 	// 函数体
// }

// 泛型结构体
// type 结构体名[T 类型约束] struct {
// 	// 字段	T
// }

// 泛型接口
// type 接口名[T 类型约束] interface {
// 	// 方法
// }

// 为什么用泛型
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 泛型方式：一个函数，可以适配不同类型
// ~符号表示包括底层类型在内的所有类型。比如，不仅包括 int，还包括基于 int 定义的新类型
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
