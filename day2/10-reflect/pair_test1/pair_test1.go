/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 17:30:37
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-02 17:31:28
 * @FilePath: \go-learning\day2\10-reflect\pair.go
 * @Description: pair 例1
 *
 */
package main

import "fmt"

func main() {
	//pair<statictype:string, value:hello>
	var a string
	a = "hello"
	var allType interface{}
	allType = a
	v, ok := allType.(string)
	fmt.Println(v, ok)

}
