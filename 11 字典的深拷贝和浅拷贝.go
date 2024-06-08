package main

// import (
// 	"fmt"
// )

// func CopyMap(s1 []string) {
// 	s3 := make([]string, len(s1), cap(s1))
// 	copy(s3, s1)
// 	var studentName map[int]string
// 	studentName = make(map[int]string)
// 	for k, v := range s3 {
// 		studentName[k+1] = v
// 	}
// 	fmt.Println("深度拷贝", studentName)
// }

// func main() {
// 	// 浅拷贝
// 	teachename := map[int]string{
// 		1: "算啦上次",
// 		2: "脚手架",
// 		3: "刹车",
// 		4: "jjjf",
// 	}
// 	// studentName := teachename
// 	// fmt.Println("teachename", teachename)
// 	// fmt.Println("studentName", studentName)
// 	// 深拷贝 -- 手动进行深拷贝
// 	var s2 = []string{
// 		"sjjsl",
// 	}
// 	for i := 1; i <= 4; i++ {
// 		s2 = append(s2, teachename[i])
// 	}
// 	s2 = s2[1:]
// 	// fmt.Println("s2", s2)
// 	var s4 []string
// 	s4 = make([]string, len(s2), cap(s2))
// 	copy(s4, s2)
// 	fmt.Println(s4)
// 	CopyMap(s4)
// }
