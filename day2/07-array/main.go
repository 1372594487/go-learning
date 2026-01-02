/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-17 20:50:48
 * @LastEditors: 1372594487 1372594487@qq.com
 * @LastEditTime: 2025-12-17 21:55:33
 * @FilePath: \go-learning\day2\07-array\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

 package main
 import  "fmt"
 type CartItem struct {
     Name string
     Price float64
	 Quantity int 
 }
func main()  {
	var cart []CartItem
	// cart:= make([]CartItem,0)

	cart = append(cart,CartItem{Name:"手机",Price:3999.99,Quantity:1})
	cart = append(cart,CartItem{Name:"电脑",Price:5999.99,Quantity:1})
	cart = append(cart,CartItem{Name:"电视",Price:4999.99,Quantity:1})

	total := 0.0
	for _,item := range cart {
	    fmt.Printf("商品：%s，单价：%.2f，数量：%d\n",item.Name,item.Price,item.Quantity)
		total += item.Price * float64(item.Quantity)
	}

	fmt.Printf("总价：%.2f\n",total)
	fmt.Printf("商品数量：%d\n",len(cart))
}