/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-24 21:14:55
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-25 20:51:20
 * @Description: File description
 */
package user

import (
	"12-package/common"
	"fmt"
	"time"
)

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

func (u *User) DisplayInfo() error {
	fmt.Printf("用户ID：%d,用户名：%s,邮箱：%s\n", u.ID, u.Username, u.Email)
	return nil
}

func (u *User) ValidatePassword() bool {
	return len(u.Password) >= 6
}

func CreateUser(name, email, password string) *User {
	common.LogOperation("创建用户")
	user := &User{
		ID:       time.Now().Unix(),
		Username: name,
		Email:    email,
		Password: password,
	}
	if !user.ValidatePassword() {
		common.LogOperation("密码验证失败，用户创建失败")
		return user
	}

	return user
}
