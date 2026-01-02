/*
* @Author: 1372594487 1372594487@qq.com
* @Date: 2025-12-28 15:08:52
  - @LastEditors: '1372594487
  - @LastEditTime: 2025-12-28 16:30:12

* @Description: 字符串处理（strings/split/join/trim/regexp...）
Go语言的strings包提供了丰富的字符串操作函数，
strconv包提供了字符串和基本数据类型之间的转换，
regexp提供了正则表达式的功能
应用场景：
string包：用户输入验证、文本解析、数据清洗、URL处理
srrconv包：配置文件数值转换、命令行参数转换、数据库数据转换
regexp包：邮箱/手机号验证、复杂文本提取、数据格式标准化、日志分析
*/
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// validateEmail 验证邮箱格式的函数
// 参数:
//
//	email: 需要验证的邮箱字符串
//
// 返回值:
//
//	bool: 如果邮箱格式正确返回true，否则返回false
func validateEmail(email string) bool {
	// TODO: 实现邮箱验证逻辑
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func validatePhoneNumber(phoneNumber string) bool {
	pattern := `^\+861[3456789]\d{9}$`
	match, _ := regexp.MatchString(pattern, phoneNumber)
	return match
}

func processUserInput(input string) map[string]interface{} {
	result := make(map[string]interface{})
	cleaned := strings.TrimSpace(input)
	cleaned = strings.ToLower(cleaned)
	fields := strings.Split(cleaned, ":")
	if len(fields) <= 0 {
		result["error"] = "Invalid input format"
		return result
	}
	result["name"] = fields[0]
	age, err := strconv.Atoi(strings.TrimSpace(fields[1]))
	if err != nil {
		result["age"] = "Invalid age format"
	} else {
		result["age"] = age
	}

	email := strings.TrimSpace(fields[2])
	if !validateEmail(email) {
		result["email"] = "Invalid email format"
	} else {
		result["email"] = email
	}

	phoneNumber := strings.TrimSpace(fields[3])
	if !validatePhoneNumber(phoneNumber) {
		result["phone_number"] = "Invalid phone number format"
	} else {
		result["phone_number"] = phoneNumber
	}

	result["origin_length"] = len(input)
	result["cleaned_length"] = len(cleaned)
	result["rune_count"] = utf8.RuneCountInString(input)
	return result
}

func main() {
	userInput := []string{
		"Alice:30:alice@example.com:+8613800138000",
		"Bob:25:bob@example:com:+8613800138000",
		"Charlie:35:charlie@example.com:+8613800138000",
	}
	for _, input := range userInput {
		result := processUserInput(input)
		fmt.Println("=========================")
		for key, value := range result {
			if key == "error" {
				fmt.Printf("Invalid input: %s\n", key, value)
				println("Error:", value)
			} else {
				fmt.Printf("%s: %v\n", key, value)
			}
		}
	}

	//字符串构造器
	var builder strings.Builder
	builder.WriteString("Hello, ")
	products := []string{"apple", "banana", "cherry"}
	for _, v := range products {
		builder.WriteString(v)
		builder.WriteString("\n")
	}
	fmt.Println(builder.String())

	//字符串替换和重复
	template := "通知：{message} | 重复提示：{repeat}"
	replaced := strings.Replace(template, "{message}", "系统通知", 1)
	repeated := strings.Repeat("注意！", 3)
	final := strings.Replace(replaced, "{repeat}", repeated, 1)
	fmt.Println(final)

	// 字符串前后缀
	url1 := "https://www.example.com"
	url2 := "https://www.example.com/api/v1/products"

	if strings.HasPrefix(url1, "https://") && strings.HasSuffix(url1, ".com") {
		fmt.Println("URL1 is valid")
	}
	if strings.HasPrefix(url2, "https://") && strings.HasSuffix(url2, ".com") {
		fmt.Println("URL2 is valid")
	}
}
