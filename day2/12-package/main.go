package main

import (
	"12-package/product"
	"12-package/user"
	"fmt"
)

func main() {
	userService := user.NewUserService()
	productService := product.NewProductService()

	fmt.Println("用户服务器测试")
	userService.AddUser(user.CreateUser("张三", "123456", "zhangsan@163.com"))
	userService.AddUser(user.CreateUser("李四", "123456", "lisi@163.com"))

	userService.ListUsers()

	fmt.Println("商品服务器测试")
	findID := productService.AddProduct("苹果", 5.99, 100)
	productService.AddProduct("香蕉", 2.99, 200)

	productService.FindProductById(findID).DisplayInfo()
}
