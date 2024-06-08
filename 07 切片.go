package main

// import "fmt"

// func main() {
// 	// 切片和数组是不同的 -- 切片的长度是不固定的
// 	// 定义一个切片 -- var 切片名 []类型
// 	var s1 = []int{}
// 	fmt.Println("切片的初始数据为：", s1)
// 	// 当你声明了一个切片是会有两个默认属性 -- 长度和容量
// 	fmt.Println("切片的初始长度为：", len(s1))
// 	fmt.Println("切片的初始容量为：", cap(s1))
// 	// 如何向切片中添加元素 -- append方法
// 	s1 = append(s1, 1, 2, 3, 3)
// 	fmt.Println("切片的数据为：", s1)
// 	fmt.Println("切片的长度为：", len(s1))
// 	fmt.Println("切片的容量为：", cap(s1))
// 	// 切片的第二种方法 -- make方法
// 	s2 := make([]string, 5, 10)
// 	fmt.Println("切片的初始数据为：", s2)
// 	fmt.Println("切片的初始长度为：", len(s2))
// 	fmt.Println("切片的初始容量为：", cap(s2))
// 	s2 = append(s2, "sh", "sx", "1")
// 	fmt.Println("切片的数据为：", s2)
// 	// 若添加的数据超过了切片的最大容量，切片在扩容时会成倍扩容
// 	// 切片也可以通过下标来修改值
// 	s2[0] = "sssss"
// 	fmt.Println("切片的初始数据为：", s2)
// 	// 也可以通过for循环
// 	for k, v := range s2 {
// 		fmt.Printf("第%d个数据为：%s\n", k+1, v)
// 	}
// }
