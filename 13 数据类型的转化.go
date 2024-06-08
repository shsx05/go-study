package main

// import (
// 	"fmt"
// 	"math"
// 	"reflect"
// 	"strconv"
// )

// func calculateArea(r float64) float64 {
// 	s := math.Pi * r * r
// 	return s
// }
// func main() {
// 	// int 转 float64
// 	r := 3
// 	s := calculateArea(float64(r))
// 	fmt.Println("面积为：", s)
// 	// int 转化为字符串
// 	count := 3
// 	str := strconv.Itoa(count)
// 	fmt.Println("字符串为：", str)
// 	fmt.Println("str的类型", reflect.TypeOf(str))
// 	// 字符串转化为int
// 	str1 := "3a"
// 	count1, err := strconv.Atoi(str1)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("字符串为：", count1)
// 	}
// 	// bool型
// 	s1 := "t"
// 	b1, err := strconv.ParseBool(s1)
// 	fmt.Println(b1)
// }
