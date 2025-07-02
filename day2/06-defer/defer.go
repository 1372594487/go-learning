/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-02 10:27:20
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-02 10:45:52
 * @FilePath: \go-learning\day2\06-defer\defer.go
 * @Description:defer关键字
 *
 */
package main

import "fmt"

//defer 语句会将函数推迟到外层函数返回之后执行。推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
//另外，defer 语句中的函数调用会被压入一个栈中，因此后执行的 defer 语句中的函数会先被调用。因此，如果需要在函数返回前执行多个操作，可以将这些操作放在不同的 defer 语句中，以确保它们按照正确的顺序执行。
func deferDemo() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

//输出结果：
//hello
//world

//defer 语句通常用于简化函数返回值的计算，或者在函数执行完毕后进行清理操作。例如，可以在打开文件后使用 defer 语句关闭文件，以确保文件在使用完毕后被正确关闭。
//func openFile() {
//	file, err := os.Open("file.txt")
//	if err != nil {
//		return
//	}
//	defer file.Close()
//	// do something with file
//}
//在上面的例子中，如果打开文件失败，函数会提前返回，但 defer 语句中的 file.Close() 仍然会被执行，以确保文件被正确关闭。

//另外，defer 语句还可以用于处理错误。例如，可以在打开文件后使用 defer 语句关闭文件，并在打开文件失败时返回错误。这样可以确保文件在使用完毕后被正确关闭，即使函数在打开文件失败时提前返回。
//func openFile() (*os.File, error) {
//	file, err := os.Open("file.txt")
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//	return file, nil
//}

//需要注意的是，defer 语句中的函数参数在 defer 语句执行时会被计算，而不是在函数返回时。因此，如果 defer 语句中的函数参数是一个变量，并且该变量的值在 defer
// 语句执行后会发生变化，那么 defer 语句中的函数参数的值将是 defer 语句执行时的值，而不是函数返回时的值。
func deferDemo1() {
	i := 0
	defer fmt.Println(i) // 输出 0
	i++
	defer fmt.Println(i) // 输出 1
	return
}

//在上面的例子中，defer 语句中的函数参数 i 在 defer 语句执行时会被计算，因此 defer 语句中的函数参数的值将是 defer 语句执行时的值，而不是函数返回时的值。
// 因此，第一个 defer 语句中的函数参数的值是 0，第二个 defer 语句中的函数参数的值是 1。

// defer和 return的执行顺序
func deferFunc() int {
	defer fmt.Println("deferFunc call")
	return 0
}
func returnFunc() int {
	defer fmt.Println("returnFunc call")
	return 0
}
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}
func main() {
	deferDemo()
	deferDemo1()
	returnAndDefer()

}
