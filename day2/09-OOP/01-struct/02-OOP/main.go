/*
* @Author: 1372594487 1372594487@qq.com
* @Date: 2025-12-22 23:21:26
  - @LastEditors: '1372594487
  - @LastEditTime: 2025-12-23 01:56:07

* @Description:结构体与方法的组合应用
面向对象编程（OOP）是一种编程范式，强调将数据和操作数据的行为封装在一起，通过对象来组织代码。
应用场景：
真实业务场景：如银行系统、电商系统、管理系统等。
电商平台有不同类型的用户(普通用户、VIP用户)，这些用户有红铜属性（基础信息）和不同特性
（权限、折扣等）。通过组合可以避免重复代码，同时保持类型之间的清晰关系。
*/
package main

import (
	"fmt"
	"time"
)

const (
	UserTypeNormal = "normal"
	UserTypeVIP    = "vip"
)

type BaseUser struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BaseUser) GetCreatedDate() string {
	return b.CreatedAt.Format("2006-01-02")
}

func (b *BaseUser) DisplayBasiceInfo() {
	fmt.Printf("ID:%d,用户名:%s,邮箱:%s,创建时间：%s\n", b.ID, b.Name, b.Email, b.CreatedAt)
}

type Address struct {
	Province string
	City     string
	District string
	Detail   string
}

type NormalUser struct {
	BaseUser
	Address []Address
}

func (a Address) GetFullAddress() Address {
	return a
}

func (n *NormalUser) AddAddress(address Address) {
	n.Address = append(n.Address, address)
}

type VIPUser struct {
	BaseUser
	Address    []Address
	VIPLevel   int
	Discount   float64
	ExpireTime time.Time
}

func (v *VIPUser) IsVipValid() bool {
	return v.ExpireTime.After(time.Now())
}

func (v *VIPUser) GetDiscount() float64 {
	if v.IsVipValid() {
		fmt.Printf("VIP用户享受折扣:%.2f\n", v.Discount)
		return v.Discount
	} else {
		fmt.Println("VIP用户已过期")
	}
	return 1.0
}

type UserService struct {
}

func genUserID() int64 {
	return time.Now().Unix()
}

func (u *UserService) RegisterUser(name, email, password string, userType string) (interface{}, error) {
	base := BaseUser{
		ID:        genUserID(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
	switch userType {
	case UserTypeNormal:
		// TODO: 注册普通用户
		normal := NormalUser{
			BaseUser: base,
			Address:  []Address{},
		}
		return normal, nil
	case UserTypeVIP:
		// TODO: 注册VIP用户
		vip := VIPUser{
			BaseUser:   base,
			VIPLevel:   1, // 默认VIP等级为1
			Discount:   0.9,
			ExpireTime: time.Now().AddDate(0, 0, 30),
		}
		return vip, nil
	default:
		return 0, fmt.Errorf("未知的用户类型:%s", userType)
	}
}
func main() {
	service := &UserService{}
	normalUser, err := service.RegisterUser("张三", "zhangsan@example.com", "password123", UserTypeNormal)
	if err != nil {
		fmt.Printf("注册普通用户失败: %v\n", err)
		return
	}
	fmt.Printf("注册普通用户成功，ID: %d\n", normalUser)

	vipUser, err := service.RegisterUser("李四", "lisi@example.com", "password456", UserTypeVIP)
	if err != nil {
		fmt.Printf("注册VIP用户失败: %v\n", err)
		return
	}
	fmt.Printf("注册VIP用户成功，ID: %d\n", vipUser)

	nUser := normalUser.(NormalUser)
	nUser.AddAddress(Address{Province: "广东省", City: "广州市", District: "天河区", Detail: "XX街道XX号"})
	nUser.DisplayBasiceInfo()

	users := []interface{}{nUser, vipUser}
	fmt.Println("所有注册用户信息:", users)
	// for _, user := range users {
	// 	switch u := user.(type) {
	// 	case NormalUser:
	// 		fmt.Println("普通用户:")
	// 		fmt.Println("用户ID：", u.ID)
	// 		fmt.Println("用户名：", u.Name)
	// 		fmt.Println("邮箱：", u.Email)
	// 		fmt.Println("创建时间：", u.CreatedAt)
	// 		fmt.Println("地址：")
	// 		for _, addr := range u.Address {
	// 			fmt.Println("-", addr.GetFullAddress())
	// 		}

	// 	case VIPUser:
	// 		fmt.Println("VIP用户:")
	// 		fmt.Println("用户ID：", u.ID)
	// 		fmt.Println("用户名：", u.Name)
	// 		fmt.Println("邮箱：", u.Email)
	// 		fmt.Println("创建时间：", u.CreatedAt)
	// 		fmt.Println("VIP等级：", u.VIPLevel)
	// 		fmt.Println("折扣：", u.GetDiscount())
	// 		fmt.Println("VIP状态", u.IsVipValid())
	// 		fmt.Println("地址：")
	// 		for _, addr := range u.Address {
	// 			fmt.Println("-", addr.GetFullAddress())
	// 		}

	// 	default:
	// 		fmt.Println("未知用户类型")
	// 	}
	// }

	//notify.go 文件中的接口与结构体组合应用
	notifyDemo()
}
