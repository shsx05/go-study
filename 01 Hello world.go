// package: 声明这个文件是属于那个包的
// 包的定义：可以理解为go源码的集合，也是一种高级的代码复用方案

// main是一个特殊的包，一个可执行的程序有且只有一个main包
// main函数是整个程序的入口
package main

// import: 用来导入包
import "fmt"

// 定义一个入口函数
func main() {
	fmt.Println("Hello World! hello go")
}
