package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	// 定义变量-第一种方法
// 	var name string = "shsx"
// 	fmt.Println("我的名字是：", name)
// 	var name2 string
// 	name2 = "shihao"
// 	fmt.Println("name2：", name2)
// 	var (
// 		name3 string = "name3"
// 		name4 int    = 10
// 	)
// 	fmt.Println(name3, name4)
// 	// 第二种方法 --- 语法糖的形式
// 	b := true
// 	f := 12345.4432
// 	fmt.Println(b, f)
// 	// reflect来查看变量的类型
// 	fmt.Println("b的类型: ", reflect.TypeOf(b))

// 	// 定义常量
// 	const c1, c2, c3 = 1, 2, 3
// 	fmt.Println(c1, c2, c3)
// 	// 特殊常量 --- iota
// 	// 可以被编译器修改的常量
// 	// 当iota在关键子const出现之前会被重置为0，当const定义多变量是每多一个变量，iota加一
// 	const (
// 		d1 = iota
// 		d2 = "shsx"
// 		d3 = iota
// 	)
// 	fmt.Println(d1, d2, d3)
// }
