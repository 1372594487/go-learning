/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-24 22:41:44
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-25 22:30:20
 * @Description: File description
 */
package product

import (
	"12-package/common"
	"fmt"
	"log"
	"time"
)

type product struct {
	ID    int64
	Name  string
	Price float64
	Stock int
}

type productService struct {
	products []*product
	logger   *log.Logger
}

func NewProductService() *productService {
	return &productService{
		products: make([]*product, 0),
		logger:   common.NewLogger("product_service.log"),
	}
}

func (p *product) DisplayInfo() {
	fmt.Printf("商品ID：%d,名称：%s,价格：%.2f,库存：%d\n", p.ID, p.Name, p.Price, p.Stock)
}

func (ps *productService) AddProduct(name string, price float64, stock int) int64 {
	product := &product{
		ID:    time.Now().Unix(),
		Name:  name,
		Price: price,
		Stock: stock,
	}
	ps.products = append(ps.products, product)
	ps.logger.Printf("添加商品成功： %s 价格： %.2f 库存： %d\n", product.Name, product.Price, product.Stock)
	return product.ID
}

func (ps *productService) ListProducts() {
	for _, product := range ps.products {
		ps.logger.Printf("商品ID： %d 名称： %s 价格： %.2f 库存： %d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func (ps *productService) FindProductById(id int64) *product {
	ps.logger.Printf("查找商品ID： %d\n", id)
	for _, product := range ps.products {
		if product.ID == id {
			ps.logger.Printf("找到商品： %s\n", product.Name)
			return product
		}
	}
	ps.logger.Printf("未找到商品ID： %d\n", id)
	return nil
}
