/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-30 14:35:14
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-30 18:13:42
 * @Description: File description
 */
package main

import (
	"cmp"
	"fmt"
)

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v, "")
	}
	fmt.Println()
}

type Stringer interface {
	String() string
}

func PrintAll[T Stringer](items []T) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Sum[T Number](numbers ...T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PrintSlice(intSlice)
	fmt.Println(FindIndex(intSlice, 5)) // 4

	stringSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	PrintSlice(stringSlice)

	fmt.Println(Max(1, 2))
	fmt.Println(Max("a", "b"))
	fmt.Println(Max(1.2, 2.3))

	persons := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	PrintAll(persons)

	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Sum(ints...))

	// floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
	// fmt.Println(Sum(floats...)) // Error: float64 does not satisfy Number
	// (float64 missing in ~int | ~int8 | ~int16 | ~int32 | ~int64)

	fmt.Println("======================================================")
	// go run main.go tools.go
	// TestMyTools()

	pair1 := Pair[string, int]{"hello", 42}
	// pair2 := Pair[string, int]{"hello", 42}
	pair2 := Pair[int, string]{42, "hello"}
	fmt.Println(pair1, pair2) // true
}
