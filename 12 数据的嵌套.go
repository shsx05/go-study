package main

// import (
// 	"fmt"
// )

// func main() {
// 	// order1 := map[string]int{
// 	// 	"米饭": 1,
// 	// 	"面包": 2,
// 	// 	"鸡蛋": 3,
// 	// }
// 	// order2 := map[string]int{
// 	// 	"炒肉":  12,
// 	// 	"炒鸡":  34,
// 	// 	"炒鸡蛋": 6,
// 	// }
// 	// order3 := map[string]int{
// 	// 	"可乐":  3,
// 	// 	"雪碧":  3,
// 	// 	"矿泉水": 1,
// 	// }

// 	// // 切片嵌套对象
// 	// var slice = []map[string]int{
// 	// 	order1,
// 	// 	order2,
// 	// 	order3,
// 	// }
// 	// // 遍历组合数据
// 	// for k, v := range slice {
// 	// 	fmt.Printf("第%d天的菜单:\n", k+1)
// 	// 	for name, price := range v {
// 	// 		fmt.Printf("\t菜名:%s, 价格:%d\n", name, price)
// 	// 	}
// 	// }

// 	// 对象嵌套对象
// 	message1 := map[string]int{
// 		"小明": 12,
// 		"小红": 13,
// 		"小黄": 12,
// 	}
// 	message2 := map[string]int{
// 		"djls":    34,
// 		"schsk":   23,
// 		"scjascx": 45,
// 	}
// 	message3 := map[string]int{
// 		"话随处可见": 23,
// 		"上次几块钱": 32,
// 		"今天几块钱": 24,
// 	}
// 	var message map[string]map[string]int
// 	message = make(map[string]map[string]int)
// 	message["大一"] = message1
// 	message["大二"] = message2
// 	message["大三"] = message3
// 	// 循环遍历
// 	for k, v := range message {
// 		fmt.Printf("%s:\n", k)
// 		for name, age := range v {
// 			fmt.Printf("\tname:%s, age:%d\n", name, age)
// 		}
// 	}
// }
