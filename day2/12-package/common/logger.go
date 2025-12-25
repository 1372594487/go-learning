/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-24 20:19:39
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-25 22:49:47
 * @Description: 包（package）与导入(import)
 应用场景：
 真实业务场景：电商微服务项目结构
 不同功能模块拆分为独立包
 user包：用户注册、登录、资料管理
 product包：商品浏览、搜索、详情
 order包：订单创建、支付、查询
 payment包：支付网关集成、退款处理
 common包：公共工具函数、常量定义
 通过包的划分，实现代码模块化，提升可维护性和复用性
*/

// 包是多个Go源码的集合，每个源码文件都属于某个包，包通过包名进行标识。
// 包的主要作用是组织代码，避免命名冲突，同时提供模块化、可重用的代码单元。
// 在Go语言中，包的命名通常使用小写字母，并且不使用下划线。
// 包声明在第一行

// package 包名
package common

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 导入

// import (
// 	"标准库包"
// 	"第三方包"
// 	"项目内部包"
// )

//单行导入： import "fmt"
// 多行导入
// import (
// 	"fmt"
// 	"math"
// )

// 可见性规则
// 大写字母开头：导出（public），包外可见，
// 小写字母开头：未导出（private），包外不可见。

func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix+"", log.Ldate|log.Ltime)
}

func LogOperation(operation string) {
	fmt.Printf("[%s]操作： %s\n", time.Now().Format(time.DateTime), operation)
}

func formatLogMessage(level, message string) string {
	return fmt.Sprintf("[%s ] %s: %s", time.Now().Format(time.DateTime), level, message)
}
