/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-20 04:21:52
 * @LastEditors: 1372594487 1372594487@qq.com
 * @LastEditTime: 2025-12-20 07:12:18
 * @FilePath: \go-learning\day2\11-Go标准库\sort包\main.go
 * @Description:
 应用场景
 1、数据展示：表格数据排序（按价格、日期、销量等）
 2、搜索排名：搜索结果按相关性排序
 3、数据分析：数据预处理和统计分析
 4、缓存优化：LRU缓存淘汰策略
 5、调度系统：任务按优先级排序执行
 */

 package main

import (
	"time"
	"sort"
	"fmt"
)

//排序基本类型
// sort.Ints([]int)
// sort.float64(slice []float64)
// sort.Strings(slice []string)



type Product struct {
    ID int 
	Name string
	Price float64
	Rating float64
	Sales int
	CreatedAt time.Time
}

type ProductManager struct {
    Products []Product
}
func NewProductManager() *ProductManager {
	return &ProductManager{
		Products: []Product{
		{ID: 1, Name: "Product 1", Price: 100.0, Rating: 4.5, Sales: 10, CreatedAt: time.Now().Add(-24 * time.Hour)},
		{ID: 2, Name: "Product 2", Price: 200.0, Rating: 4.0, Sales: 20, CreatedAt: time.Now().Add(-48 * time.Hour)},
		{ID: 3, Name: "Product 3", Price: 150.0, Rating: 4.8, Sales: 15, CreatedAt: time.Now().Add(-12 * time.Hour)},
		{ID: 4, Name: "Product 4", Price: 300.0, Rating: 4.2, Sales: 30, CreatedAt: time.Now().Add(-72 * time.Hour)},
		{ID: 5, Name: "Product 5", Price: 250.0, Rating: 4.7, Sales: 25, CreatedAt: time.Now().Add(-36 * time.Hour)},
	},
}
}

func (pm *ProductManager) SortByPrice(ascending bool) {
	if ascending {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].Price < pm.Products[j].Price
		})
	} else {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].Price > pm.Products[j].Price
		})
}
}
func (pm *ProductManager) SortByRating(ascending bool) {
	if ascending {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].Rating < pm.Products[j].Rating
		})
	} else {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].Rating > pm.Products[j].Rating
		})
	}

}
func (pm *ProductManager) SortBySales(ascending bool) {
	if ascending {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].Sales < pm.Products[j].Sales
		})
	} else {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].Sales > pm.Products[j].Sales
		})
	}

}

func (pm *ProductManager) SortByCreatedAt(ascending bool) {
	if ascending {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].CreatedAt.Before(pm.Products[j].CreatedAt)
		})
	} else {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].CreatedAt.After(pm.Products[j].CreatedAt)
		})
	}

}

func (pm *ProductManager) SortByMutiple(product Product) {
	sort.Slice(pm.Products, func(i, j int) bool {
		// 比较价格
		if pm.Products[i].Price != pm.Products[j].Price {
			return pm.Products[i].Price < pm.Products[j].Price
		}
		// 价格相同，比较评分 
		 if pm.Products[i].Rating != pm.Products[j].Rating {
			return pm.Products[i].Rating < pm.Products[j].Rating
		} 
		// 价格和评分都相同，比较销量
		 if pm.Products[i].Sales != pm.Products[j].Sales {
			return pm.Products[i].Sales < pm.Products[j].Sales
		} 
		// 前面都相同，比较创建时间
			return pm.Products[i].CreatedAt.Before(pm.Products[j].CreatedAt)
	})
}

func (pm *ProductManager) PrintProducts() {
	for _, product := range pm.Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Rating: %.1f, Sales: %d, CreatedAt: %s\n", product.ID, product.Name, product.Price, product.Rating, product.Sales, product.CreatedAt)
	}
}

//排序自定义类型
type Person struct {
	Name string
	Age int
}
type People []Person



//实现sort.Interface接口的三个方法即可使用sort包中的方法进行排序
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }
// sort.Sort(slice)

func (p People) Less (i,j int) bool {
	return p[i].Age < p[j].Age
}
func (p People) Swap (i, j int) {
	p[i], p[j] = p[j], p[i]
}
// 为People类型实现sort.Interface接口
func (p People) Len() int {
    return len(p)
}


// 更灵活的sort.Slice复用方式
// 可以创建通用的排序函数
type SortConfig struct {
    Field     string
    Ascending bool
}

func SortPeople(people []Person, config SortConfig) {
    sort.Slice(people, func(i, j int) bool {
        var less bool
        switch config.Field {
        case "age":
            less = people[i].Age < people[j].Age
        case "name":
            less = people[i].Name < people[j].Name
        }
        return config.Ascending == less
    })
}

func main() {

	
	pm := NewProductManager()

	fmt.Println("-----Products before sorting:-----")
	pm.PrintProducts()

	pm.SortBySales(true)
	fmt.Println("-----Products after sorting by sales (descending):-----")
	pm.PrintProducts()

	pm.SortByCreatedAt(true)
	fmt.Println("-----Products after sorting by created at (ascending):-----")
	pm.PrintProducts()

	pm.SortByMutiple(Product{ID: 1, Name: "Product 1", Price: 100.0, Rating: 4.5, Sales: 10, CreatedAt: time.Now()})
	fmt.Println("-----Products after sorting by multiple fields:-----")
	pm.PrintProducts()

	    people := People{
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 20},
		{"David", 35},
    }
    fmt.Println("-----People after sorting:-----")
    sort.Sort(people)
    for _, person := range people {
        fmt.Printf("Name:%s, Age:%d\n", person.Name, person.Age)
    }
	names := []string{"Alice", "Bob", "Charlie", "David"}
	sort.Strings(names)
	index := sort.SearchStrings(names, "Charlie")
	fmt.Println("Charlie's index is",index) // 输出: 2

	 // 同一个函数实现不同的排序需求
	fmt.Println("-----People after sorting by age (descending):-----")
    SortPeople(people, SortConfig{Field: "age", Ascending: false})
	 for _, person := range people {
        fmt.Printf("Name:%s, Age:%d\n", person.Name, person.Age)
    }
	fmt.Println("-----Products after sorting by name (descending):-----")
    SortPeople(people, SortConfig{Field: "name", Ascending: false})	
	 for _, person := range people {
        fmt.Printf("Name:%s, Age:%d\n", person.Name, person.Age)
    }
	

}