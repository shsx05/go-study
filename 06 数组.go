package main

// import "fmt"

// func main() {
// 	// 定义一个数组
// 	// var 数组名 = [数组的长度]数组的类型{数据}--1
// 	var className = [3]string{"sh", "sx", "sq"}
// 	fmt.Println(className)
// 	// 用语法糖形式
// 	classAge := [3]int{10, 13, 16}
// 	fmt.Println(classAge)
// 	// 获取数组中的数据 -- 数组名[下标]
// 	classname1 := className[0]
// 	classage1 := classAge[0]
// 	fmt.Println("第一名的姓名：", classname1)
// 	fmt.Println("第一名的年龄：", classage1)
// 	// 获取数组中的全部数据 -- for循环
// 	// 获取数组的长度 -- len()
// 	for i := 0; i < len(className); i++ {
// 		fmt.Printf("第%d个老师的姓名:%s\n", i+1, className[i])
// 	}
// 	fmt.Println("************************************")
// 	// 第二种方法 -- range
// 	for index, age := range classAge {
// 		fmt.Printf("第%d位学生的年龄:%d\n", index+1, age)
// 	}
// 	// 自动判断长度
// 	classSroce := [...]int{1, 2, 3, 4, 5, 6, 7, 6, 4, 2, 1, 7}
// 	fmt.Println("数组的长度为：", len(classSroce))
// }
