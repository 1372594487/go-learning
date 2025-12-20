/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 16:00:47
 * @LastEditors: 1372594487 1372594487@qq.com
 * @LastEditTime: 2025-12-20 18:22:49
 * @FilePath: \go-learning\day2\08-struct\struct.go
 * @Description: struct定义与使用
 应用场景：
1、用户管理系统：在电商平台、社交应用等系统中，需要管理用户信息，如用户名、密码、年龄、性别、地址等。
可以使用结构体来表示用户信息，每个用户都是一个结构体实例，包含用户的各种属性。
2、
 *
 */
package main

import "fmt"

//结构体是Go语言中组织相关数据的复合类型，将多个不同类型的数据组合在一起形成的数据类型
// 结构体是值类型，在函数间传递时，会复制一份新的结构体，而不是传递指针

// 结构体初始化方式
// var p Person //各字段为零值
// person :=Person{name: "张三", age: 18} //结构体字面量
// new(Person) //返回指针


// type关键字定义结构体
type Person struct {
	name string
	age  int
}


type User struct{
	ID int
	UserName string
	Password string
	Email string
	Age int
	Address Address
}

type Address struct{
	Province string
	City string
	Street string
	ZipCode string
	Detail string
}

type UserProfile struct {
	UserName string
	Age int
	Email string
	Address
}

func main() {
	addr := Address{Province: "北京市", City: "北京市", Street: "朝阳区", ZipCode: "100000", Detail: "xx街道xx号"}
	user1 := User{ID: 1, UserName: "张三", Password: "123456", Email: "123456@qq.com", Age: 18, Address: addr}
	user2 := User{ID: 2, UserName: "李四", Password: "123456", Email: "123456@qq.com", Age: 18, Address: Address{
		Province: "广西省", City: "南宁市", Street: "武鸣区", ZipCode: "000000", Detail: "xx街道xx号",
	}}
	user3 := &User{ID: 3, UserName: "王五", Password: "123456", Email: "123456@qq.com", Age: 18, Address: Address{
	 Province: "广东省", City: "广州市", Street: "天河区", ZipCode: "000000", Detail: "xx街道xx号",
	}
	}

	UserProfile := UserProfile{UserName: "赵六", Age: 18, Email: "123456@qq.com", Address: Address{
		Province: "广东省", City: "广州市", Street: "天河区", ZipCode: "000000", Detail: "xx街道xx号",
	}}

	fmt.Printf("User1: %+v\n,City: %s\n,Street: %s\n", user1, user1.Address.City, user1.Address.Street)
	fmt.Printf("User1: %+v\n,City: %s\n,Street: %s\n", user2, user2.Address.City, user2.Address.Street)
	fmt.Printf("User3: %+v\n,City: %s\n,Street: %s\n",user3, user3.Address.City, user3.Address.Street)
	fmt,Printf("UserProfile: %+v\n,City: %s\n,Street: %s\n",UserProfile, UserProfile.City, UserProfile.Street)
}
