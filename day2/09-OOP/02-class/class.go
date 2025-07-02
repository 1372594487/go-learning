/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 16:19:30
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-02 16:37:24
 * @FilePath: \go-learning\day2\08-OOP\class\class.go
 * @Description:Go中类的表示 封装 继承
 *
 */
package main

import "fmt"

// class
type Hero struct {
	Name  string
	Ad    int
	level int
}

func (this *Hero) GetName() string {
	return this.Name
}
func (this *Hero) SetName(name string) {
	this.Name = name
}
func (this *Hero) String() string {
	return fmt.Sprintf("Name: %s, Ad: %d, level: %d", this.Name, this.Ad, this.level)
}
func (this *Hero) Walk() {
	fmt.Println(this.Name, "is walking fast")
}

/* ----------------------------------------------------------------------- */
//类的继承

type Student struct {
	Hero
	Age int
}

// 重写父类的方法
func (this *Student) Walk() {
	fmt.Println(this.Name, "is walking slow")
}

// 子类的新方法
func (this *Student) Study() {
	fmt.Println(this.Name, "is studying")
}
func (this *Student) String() string {
	return fmt.Sprintf("Name: %s, Ad: %d, level: %d, Age: %d", this.Name, this.Ad, this.level, this.Age)
}

/* ----------------------------------------------------------------------- */
//类的多态
type Bird struct {
	Name string
}

func (this *Bird) Walk() {
	fmt.Println(this.Name, "is walking")
}

func (this *Bird) Fly() {
	fmt.Println(this.Name, "is flying")
}

type Plane struct {
	Name string
}

func main() {
	h := Hero{Name: "张三", Ad: 100, level: 1}
	fmt.Println(h.GetName())
	h.SetName("李四")
	fmt.Println(h.GetName())
	fmt.Println(h.String())
	fmt.Println("-------------------------")
	// 定义一个子类对象
	var s Student
	s.Name = "张三"
	s.Ad = 100
	s.level = 1
	s.Age = 18
	fmt.Println(s.String())
	h.Walk()
	s.Walk()
	s.Study()
	fmt.Println("-------------------------")

}
